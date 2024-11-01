package workspace

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/pkg/errors"
)

type fnIsNodeReady func(ActionNodeModel) bool

/* -----------------------------------------------------------------------
   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
                               Workspace Retrieval
   ----------------------------------------------------------------------- */

func (m Model) Get(ctx context.Context) ([]WorkspaceModel, error) {
	workspaces := make([]WorkspaceModel, 0)
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
	workspace.Nodes = make([]ActionNodeModel, 0)
	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&workspace)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrWorkspaceNotFound, err)
	}
	return &workspace, nil
}

func (m Model) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]WorkspaceModel, error) {
	workspaces := make([]WorkspaceModel, 0)

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

func (m Model) GetByActionId(ctx context.Context, actionId primitive.ObjectID) ([]WorkspaceModel, error) {
	workspaces := make([]WorkspaceModel, 0)

	filter := bson.M{
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": actionId,
			},
		},
	}

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

/* -----------------------------------------------------------------------
   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
                               Node and Workspace Modification
   ----------------------------------------------------------------------- */

func (m Model) updateNodesStatus(userId primitive.ObjectID, actionId primitive.ObjectID, status string) error {
	filter := bson.M{
		"user_id": userId,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": actionId,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"nodes.$.status": status,
		},
	}

	res, err := m.Collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.ErrWorkspaceNotFound
	}

	return nil
}

func (m Model) updateNodeById(ctx context.Context, workspaceId primitive.ObjectID, node UpdateActionNodeModel) (*WorkspaceModel, error) {
	filter := bson.M{
		"_id": workspaceId,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"node_id": node.NodeId,
			},
		},
	}

	nodeUpdate := bson.M{
		"$set": bson.M{
			"nodes.$.parents":  node.Parents,
			"nodes.$.children": node.Children,
			"nodes.$.x_pos":    node.XPos,
			"nodes.$.y_pos":    node.YPos,
		},
	}

	if node.ActionId != nil {
		nodeUpdate["$set"] = bson.M{"nodes.$.action_id": node.ActionId}
	}

	for k, v := range node.Input {
		nodeUpdate["$set"].(bson.M)[fmt.Sprintf("nodes.$.input.%s", k)] = v
	}

	// Execute the update operation for the current node
	res, err := m.Collection.UpdateOne(ctx, filter, nodeUpdate)
	if err != nil {
		return nil, fmt.Errorf("error updating workspace node: %w", err)
	}
	if res.MatchedCount == 0 {
		return nil, errors.ErrNodeNotFound
	}

	var workspace WorkspaceModel

	err = m.Collection.FindOne(ctx, filter).Decode(&workspace)

	if err != nil {
		return nil, errors.ErrWorkspaceNotFound
	}

	return &workspace, nil
}

func (m Model) UpdateById(ctx context.Context, workspaceId primitive.ObjectID, update *UpdateWorkspaceModel) (*WorkspaceModel, error) {

	for _, node := range update.Nodes {
		_, err := m.updateNodeById(ctx, workspaceId, node)

		if err == nil {
			continue
		}

		if err != errors.ErrNodeNotFound {
			// Something went wrong; node could not be updated nor created
			log.Printf("Node could not be updated: %s\n", err)
			continue
		}
		_, err = m.AddNode(ctx, workspaceId, node)

		if err != nil {
			log.Printf("Node could not be created: %s\n", err)
		}
	}

	var workspace WorkspaceModel
	// Retrieve the updated workspace document
	err := m.Collection.FindOne(ctx, bson.M{"_id": workspaceId}).Decode(&workspace)
	if err != nil {
		return nil, fmt.Errorf("error retrieving updated workspace: %w", err)
	}

	return &workspace, nil
}

func (m Model) AddNode(ctx context.Context, workspaceId primitive.ObjectID, node UpdateActionNodeModel) (*WorkspaceModel, error) {
	filter := bson.M{
		"_id": workspaceId,
	}

	if node.ActionId == nil {
		return nil, fmt.Errorf("cannot create node if actionId is nil")
	}

	addNode := ActionNodeModel{
		NodeId:   node.NodeId,
		Input:    node.Input,
		Output:   make(map[string]string),
		ActionId: *node.ActionId,
		Parents:  node.Parents,
		Children: node.Children,
		Status:   "inactive",
		XPos:     node.XPos,
		YPos:     node.YPos,
	}

	nodeUpdate := bson.M{
		"$push": bson.M{
			"nodes": addNode,
		},
	}

	res, err := m.Collection.UpdateOne(ctx, filter, nodeUpdate)

	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, errors.ErrCreatingNode
	}

	var workspace WorkspaceModel

	err = m.Collection.FindOne(ctx, filter).Decode(&workspace)

	if err != nil {
		return nil, err
	}

	return &workspace, nil
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
		Nodes:  make([]ActionNodeModel, 0),
	}

	for _, node := range add.Nodes {
		node := ActionNodeModel{
			NodeId:   node.NodeId,
			ActionId: node.ActionId,
			Input:    node.Input,
			Output:   make(map[string]string),
			Parents:  node.Parents,
			Status:   "inactive",
			Children: node.Children,
			XPos:     node.XPos,
			YPos:     node.YPos,
		}
		newWorkspace.Nodes = append(newWorkspace.Nodes, node)
	}
	// Insert the blank workspace
	_, err = m.Collection.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, errors.ErrCreatingWorkspace
	}

	// Initialize the workspace
	err = m.initWorkspace(&newWorkspace, accessToken, func(node ActionNodeModel) bool { return len(node.Parents) == 0 })
	if err != nil {
		return nil, err
	}

	// Retrieve the initialized workspace
	initializedWorkspace, err := m.GetById(ctx, newWorkspace.Id)
	if err != nil {
		return nil, err
	}
	return initializedWorkspace, nil
}

/* -----------------------------------------------------------------------
   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
                               Action initialization
   ----------------------------------------------------------------------- */

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

func initAction(actionNode ActionNodeModel, accessToken string) error {

	action, _, err := action.GetByIdRequest(accessToken, actionNode.ActionId.Hex())

	if err != nil {
		return err
	}

	status, err := StartActionRequest(accessToken, actionNode, *action)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.ErrSettingAction
	}

	return nil
}

func (m Model) initWorkspace(workspace *WorkspaceModel, accessToken string, isNodeReady fnIsNodeReady) error {
	for _, node := range workspace.Nodes {
		if isNodeReady(node) {
			assignInputToAction(&node, workspace.Nodes)
			err := initAction(node, accessToken)
			if err != nil {
				return err
			}
			err = m.updateNodesStatus(workspace.UserId, node.ActionId, "active")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/* -----------------------------------------------------------------------
   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
                               Action -> Workspace Communication
   ----------------------------------------------------------------------- */

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

	update := bson.M{"$set": bson.M{}}

	for k, v := range watchCompleted.Input {
		update["$set"].(bson.M)[fmt.Sprintf("nodes.$.input.%s", k)] = v
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

	workspaces := make([]WorkspaceModel, 0)
	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func isNodeReadyToStart(child ActionNodeModel, workspace WorkspaceModel) bool {
	if child.Status != "inactive" {
		return false
	}

	for _, parent := range child.Parents {
		for _, workspaceNode := range workspace.Nodes {
			if workspaceNode.NodeId == parent && workspaceNode.Status != "completed" {
				return false
			}
		}
	}
	return true
}

func (m Model) processWorkspace(
	workspace WorkspaceModel,
	actionCompleted ActionCompletedModel,
	accessToken string,
	updatedWorkspaces *[]WorkspaceModel,
	mu *sync.Mutex,
) error {

	filter := bson.M{
		"_id": workspace.Id,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": actionCompleted.ActionId,
			},
		},
	}

	outputUpdate := bson.M{}
	for key, value := range actionCompleted.Output {
		outputUpdate[fmt.Sprintf("nodes.$.output.%s", key)] = value
	}

	update := bson.M{
		"$set": bson.M{
			"nodes.$.status": "completed",
		},
	}

	for k, v := range outputUpdate {
		update["$set"].(bson.M)[k] = v
	}

	result, err := m.Collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("Matched count: %d", result.MatchedCount)
		return fmt.Errorf("%w: %v ", errors.ErrUpdatingWorkspace, err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("%w: %s ", errors.ErrUpdatingWorkspace, "matched count is 0")
	}

	updatedResult, err := m.GetById(context.TODO(), workspace.Id)

	if err != nil {
		return err
	}

	m.initWorkspace(
		updatedResult,
		accessToken,
		func(node ActionNodeModel) bool {
			return isNodeReadyToStart(node, *updatedResult)
		},
	)

	initializedWorkspace, err := m.GetById(context.TODO(), workspace.Id)

	if err != nil {
		return err
	}

	mu.Lock()
	*updatedWorkspaces = append(*updatedWorkspaces, *initializedWorkspace)
	mu.Unlock()

	return nil
}

func (m Model) processUserWorkspaces(workspaces []WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) ([]WorkspaceModel, error) {
	var (
		wg                sync.WaitGroup
		mu                sync.Mutex
		updatedWorkspaces = make([]WorkspaceModel, 0)
		errChan           = make(chan error, len(workspaces))
	)

	for _, workspace := range workspaces {
		wg.Add(1)
		go func(ws WorkspaceModel) {
			defer wg.Done()
			if err := m.processWorkspace(ws, actionCompleted, accessToken, &updatedWorkspaces, &mu); err != nil {
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

	if len(userWorkspaces) == 0 {
		return nil, errors.ErrWorkspaceNotFound
	}
	// Iterate over all user workspaces and update them in case they have any actions that are completed
	updatedWorkspaces, err := m.processUserWorkspaces(userWorkspaces, actionCompleted, accessToken)
	if err != nil {
		return nil, err
	}

	return updatedWorkspaces, nil
}
