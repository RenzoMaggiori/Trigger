package workspace

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/pkg/errors"
)

func (m Model) Get(ctx context.Context) ([]WorkspaceModel, error) {
	var workspaces []WorkspaceModel
	filter := bson.M{}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (m Model) GetById(ctx context.Context, id primitive.ObjectID) (*WorkspaceModel, error) {
	var workspace WorkspaceModel
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&workspace)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrWorkspaceNotFound, err)
	}
	return &workspace, nil
}

func (m Model) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]WorkspaceModel, error) {
	var workspaces []WorkspaceModel

	filter := bson.M{"user_id": userId}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrWorkspaceNotFound, err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func initAction(actionNode ActionNodeModel, accessToken string) error {

	action, _, err := action.GetByIdRequest(accessToken, actionNode.ActionId.Hex())

	if err != nil {
		return err
	}

	_, err = StartActionRequest(accessToken, actionNode, *action)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Add(ctx context.Context, add *AddWorkspaceModel) (*WorkspaceModel, error) {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)

	if !ok {
		return nil, errors.ErrAccessTokenCtx
	}
	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return nil, err
	}

	newWorkspace := WorkspaceModel{
		Id:     primitive.NewObjectID(),
		UserId: session.UserId,
		Nodes:  []ActionNodeModel{},
	}

	for _, node := range add.Nodes {
		func() {
			isActive := len(node.Parents) == 0
			node := ActionNodeModel{
				NodeId:   node.NodeId,
				ActionId: node.ActionId,
				Input:    node.Input,
				Output:   nil,
				Parents:  node.Parents,
				Status:   "inactive",
				Children: node.Children,
				XPos:     node.XPos,
				YPos:     node.YPos,
			}
			defer func() {
				newWorkspace.Nodes = append(newWorkspace.Nodes, node)
			}()

			if !isActive {
				return
			}
			err = initAction(node, accessToken)
			if err != nil {
				log.Printf("Error while starting Action with id: %s", node.ActionId.Hex())
				return
			}
			node.Status = "active"
			log.Printf("Started Action with id: %s", node.ActionId.Hex())
		}()
	}

	if err != nil {
		return nil, err
	}

	_, err = m.Collection.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, errors.ErrCreatingWorkspace
	}

	return &newWorkspace, nil
}

func (m Model) ActionCompleted(ctx context.Context, actionCompleted ActionCompletedModel) ([]WorkspaceModel, error) {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return nil, errors.ErrAccessTokenCtx
	}

	// Get all user workspaces
	userWorkspaces, err := m.GetByUserId(ctx, actionCompleted.UserId)
	if err != nil {
		return nil, err
	}

	// Iterate over all user workspaces and update them in case they have any actions that are completed
	updatedWorkspaces, err := m.processUserWorkspaces(userWorkspaces, actionCompleted, accessToken)
	if err != nil {
		return nil, err
	}

	return updatedWorkspaces, nil
}

func (m Model) processUserWorkspaces(workspaces []WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) ([]WorkspaceModel, error) {
	var (
		wg                sync.WaitGroup
		mu                sync.Mutex
		updatedWorkspaces []WorkspaceModel
		errChan           = make(chan error, len(workspaces))
	)

	for _, workspace := range workspaces {
		wg.Add(1)
		go func(ws WorkspaceModel) {
			defer wg.Done()
			if err := m.processWorkspaces(ws, actionCompleted, accessToken, &updatedWorkspaces, &mu); err != nil {
				errChan <- err
			}
		}(workspace)
	}

	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return updatedWorkspaces, nil
}

func (m Model) processWorkspaces(
	workspace WorkspaceModel,
	actionCompleted ActionCompletedModel,
	accessToken string,
	updatedWorkspaces *[]WorkspaceModel,
	mu *sync.Mutex,
) error {
	updatedWorkspace, err := processWorkspaceActions(workspace, actionCompleted, accessToken)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": workspace.Id}
	update := bson.M{"$set": bson.M{"nodes": updatedWorkspace.Nodes}}
	updateResult, err := m.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if updateResult.MatchedCount == 0 {
		return errors.ErrWorkspaceNotFound
	}

	mu.Lock()
	*updatedWorkspaces = append(*updatedWorkspaces, *updatedWorkspace)
	mu.Unlock()

	return nil
}

func isChildNodeReady(child ActionNodeModel, workspace WorkspaceModel) bool {
	for _, parent := range child.Parents {
		for _, workspaceNode := range workspace.Nodes {
			if workspaceNode.NodeId == parent && workspaceNode.Status != "completed" {
				return false
			}
		}
	}
	return true
}

func findNodeById(workspace WorkspaceModel, nodeId string) *ActionNodeModel {
	for _, node := range workspace.Nodes {
		if nodeId == node.NodeId {
			return &node
		}
	}
	return nil
}

func processWorkspaceActions(workspace WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) (*WorkspaceModel, error) {
	for i, node := range workspace.Nodes {
		// Check if the has been completed
		if !(node.Status == "active" && node.ActionId == actionCompleted.ActionId) {
			continue
		}

		workspace.Nodes[i].Output = actionCompleted.Output
		workspace.Nodes[i].Status = "completed"

		if err := processChildNodes(&workspace, node.Children, accessToken); err != nil {
			return nil, err
		}
	}

	return &workspace, nil
}

func assignInputToAction(action *ActionNodeModel, workspaceNodes []ActionNodeModel) {
	for key, value := range action.Input {
		for _, node := range workspaceNodes {
			prefix := fmt.Sprintf("$%s$.", node.NodeId)
			if strings.Contains(value, prefix) {
				action.Input[key] = node.Output[strings.ReplaceAll(value, prefix, "")]
			}
		}
	}
}

func processChildNodes(workspace *WorkspaceModel, children []string, accessToken string) error {
	// Iterate over all children and start the action if all parents are completed
	for _, childId := range children {
		child := findNodeById(*workspace, childId)
		if child == nil {
			return errors.ErrNodeNotFound
		}

		if isChildNodeReady(*child, *workspace) {
			assignInputToAction(child, workspace.Nodes)
			err := initAction(*child, accessToken)
			if err != nil {
				return err
			}
			child.Status = "active"
		}
	}
	return nil
}

func (m Model) UpdateById(ctx context.Context, id primitive.ObjectID, update *UpdateWorkspaceModel) (*WorkspaceModel, error) {
	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrWorkspaceNotFound, err)
	}

	var updatedUserAction WorkspaceModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedUserAction)

	if err != nil {
		return nil, err
	}
	return &updatedUserAction, nil
}

func (m Model) WatchCompleted(ctx context.Context, watchCompleted WatchCompletedModel) ([]WorkspaceModel, error) {
	// Find documents with the matching UserId
	filter := bson.M{
		"user_id": watchCompleted.UserId,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": watchCompleted.ActionId,
			},
		},
	}

	// Define the update: set the output field for the matching nodes
	update := bson.M{
		"$set": bson.M{
			"nodes.$.output": watchCompleted.Output,
		},
	}

	// Perform the update
	_, err := m.Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Retrieve updated documents to return
	cursor, err := m.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var workspaces []WorkspaceModel
	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}
