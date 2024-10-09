package workspace

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("%w: %v", errWorkspaceNotFound, err)
	}
	return &workspace, nil
}

func (m Model) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]WorkspaceModel, error) {
	var workspaces []WorkspaceModel

	filter := bson.M{"user_id": userId}
	cursor, err := m.Collection.Find(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errWorkspaceNotFound, err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}

func initWorkspaceActions(workspace WorkspaceModel, accessToken string) (*WorkspaceModel, error) {

	for i, node := range workspace.Nodes {
		if node.Status == "pending" {
			action, _, err := action.GetActionByIdRequest(accessToken, node.ActionId.Hex())

			if err != nil {
				return nil, err
			}

			_, err = StartActionRequest(accessToken, node, *action)
			if err != nil {
				return nil, err
			}
			workspace.Nodes[i].Status = "active"
		}
	}
	return &workspace, nil
}

func nodeStatus(parents []string) string {
	if len(parents) == 0 {
		return "pending"
	}
	return "inactive"
}

func (m Model) Add(ctx context.Context, add *AddWorkspaceModel) (*WorkspaceModel, error) {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)

	if !ok {
		return nil, errors.ErrAccessTokenCtxKey
	}
	session, _, err := session.GetSessionByTokenRequest(accessToken)

	if err != nil {
		return nil, err
	}

	newWorkspace := WorkspaceModel{
		Id:     primitive.NewObjectID(),
		UserId: session.UserId,
		Nodes:  []ActionNodeModel{},
	}

	for _, node := range add.Nodes {
		node := ActionNodeModel{
			NodeId:   node.NodeId,
			ActionId: node.ActionId,
			Input:    node.Input,
			Output:   node.Output,
			Parents:  node.Parents,
			Children: node.Children,
			Status:   nodeStatus(node.Parents),
			XPos:     node.XPos,
			YPos:     node.YPos,
		}
		newWorkspace.Nodes = append(newWorkspace.Nodes, node)
	}

	workspacePtr, err := initWorkspaceActions(newWorkspace, accessToken)
	if err != nil {
		return nil, err
	}

	newWorkspace = *workspacePtr
	_, err = m.Collection.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, errCreatingWorkspace
	}

	return &newWorkspace, nil
}
func (m Model) ActionCompleted(ctx context.Context, updateActionCompleted ActionCompletedModel) ([]WorkspaceModel, error) {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return nil, fmt.Errorf("access token missing or invalid")
	}
	session, _, err := session.GetSessionByTokenRequest(accessToken)
	if err != nil {
		return nil, err
	}
	workspaces, err := m.GetByUserId(ctx, session.UserId)
	if err != nil {
		return nil, err
	}
	var (
		updateErr         error
		wg                sync.WaitGroup
		mu                sync.Mutex
		updatedWorkspaces []WorkspaceModel
	)

	for _, workspace := range workspaces {
		wg.Add(1)
		go func(workspace WorkspaceModel) {
			defer wg.Done()
			updatedWorkspace, err := updateWorkspaceNodes(workspace, updateActionCompleted, accessToken)
			if err != nil {
				mu.Lock()
				updateErr = err
				mu.Unlock()
				return
			}
			filter := bson.M{"_id": workspace.Id}
			update := bson.M{
				"$set": bson.M{"nodes": updatedWorkspace.Nodes},
			}
			updateResult, err := m.Collection.UpdateOne(ctx, filter, update)
			if err != nil {
				mu.Lock()
				updateErr = err
				mu.Unlock()
				return
			}
			if updateResult.MatchedCount == 0 {
				mu.Lock()
				updateErr = errWorkspaceNotFound
				mu.Unlock()
				return
			}
			mu.Lock()
			updatedWorkspaces = append(updatedWorkspaces, updatedWorkspace)
			mu.Unlock()
		}(workspace)
	}
	wg.Wait()
	if updateErr != nil {
		return nil, updateErr
	}
	return updatedWorkspaces, nil
}

func updateWorkspaceNodes(workspace WorkspaceModel, updateActionCompleted ActionCompletedModel, accessToken string) (WorkspaceModel, error) {
	for i, node := range workspace.Nodes {
		if node.Status == "active" && node.ActionId.Hex() == updateActionCompleted.Action {
			workspace.Nodes[i].Status = "completed"

			for _, child := range node.Children {
				for j := range workspace.Nodes {
					if workspace.Nodes[j].NodeId == child {
						workspace.Nodes[j].Status = "pending"
					}
				}
			}
		}
	}

	// Initialize actions (this assumes some logic to update workspace actions with accessToken)
	updatedWorkspace, err := initWorkspaceActions(workspace, accessToken)
	if err != nil {
		return workspace, err
	}

	return *updatedWorkspace, nil
}

func (m Model) UpdateById(ctx context.Context, id primitive.ObjectID, update *UpdateWorkspaceModel) (*WorkspaceModel, error) {
	filter := bson.M{"_id": id}
	updateData := bson.M{"$set": update}

	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errWorkspaceNotFound, err)
	}

	var updatedUserAction WorkspaceModel
	err = m.Collection.FindOne(ctx, filter).Decode(&updatedUserAction)

	if err != nil {
		return nil, err
	}
	return &updatedUserAction, nil

}

// func (m Model) DeleteById(ctx context.Context, id primitive.ObjectID) error {
// 	filter := bson.M{"_id": id}
// 	result, err := m.Collection.DeleteOne(ctx, filter)

// 	if err != nil {
// 		return fmt.Errorf("%w: %v", errUserActionNotFound, err)
// 	}
// 	if result.DeletedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}
// 	return nil
// }
