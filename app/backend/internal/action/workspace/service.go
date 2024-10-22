
package workspace

import (
	"context"
	"fmt"
	"log"
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

// initWorkspaceActions initializes the actions for the nodes in the given workspace.
// It iterates through the nodes in the workspace and starts actions for nodes with a "pending" status.
// If an action is successfully started, the node's status is updated to "active".
//
// Parameters:
//   - workspace: The WorkspaceModel containing the nodes to be initialized.
//   - accessToken: The access token used for authentication in action requests.
// Returns:
//   - A pointer to the updated WorkspaceModel
//   - An error if there is any error.

func initAction(actionNode ActionNodeModel, accessToken string) error {

	action, _, err := action.GetActionByIdRequest(accessToken, actionNode.ActionId.Hex())

	if err != nil {
		return err
	}

	_, err = StartActionRequest(accessToken, actionNode, *action)
	if err != nil {
		return err
	}
	return nil
}

// nodeStatus determines the status of a node based on its parent nodes.
// If the node has no parents, it returns "pending". Otherwise, it returns "inactive".
func isNodeActive(parents []string) bool {
	if len(parents) == 0 {
		return true
	}
	return false
}

// Add adds a new workspace to the collection based on the provided AddWorkspaceModel.
// initializes the workspace with the provided nodes, and inserts the new workspace into the collection.
//
// Parameters:
//   - ctx: The context for the request, which should contain the access token
//   - add: A pointer to the AddWorkspaceModel containing the details of the workspace to be added
//
// Returns:
//   - A pointer to the newly created WorkspaceModel
//   - An error if the access token is not found in the context, if there is an error fetching the session,
//     or if there is an error creating the workspace
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
		func() {
			isActive := isNodeActive(node.Parents)
			node := ActionNodeModel{
				NodeId:   node.NodeId,
				ActionId: node.ActionId,
				Input:    node.Input,
				Output:   node.Output,
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
			if err == nil {
				log.Printf("Error while starting Action with id: %s", node.ActionId.Hex())
				node.Status = "active"
				return
			}
			log.Printf("Started Action with id: %s", node.ActionId.Hex())
		}()
	}

	if err != nil {
		return nil, err
	}

	_, err = m.Collection.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, errCreatingWorkspace
	}

	return &newWorkspace, nil
}

// ActionCompleted updates the workspaces based on the provided action completion details.
//
// This method performs the following steps:
// 1. Retrieves the access token from the context.
// 2. Fetches the workspaces associated with the user ID from the updateActionCompleted model.
// 3. Concurrently processes each workspace to run actions and update the workspace nodes.
// 4. Updates the workspace in the database and collects the updated workspaces.
//
// Parameters:
// - ctx: The context for the request, which should contain the access token.
// - updateActionCompleted: The model containing the details of the action completion.
//
// Returns:
// - A slice of updated WorkspaceModel objects.
// - An error if any step in the process fails.
func (m Model) ActionCompleted(ctx context.Context, updateActionCompleted ActionCompletedModel) ([]WorkspaceModel, error) {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return nil, fmt.Errorf("access token missing or invalid")
	}
	workspaces, err := m.GetByUserId(ctx, updateActionCompleted.UserId)
	if err != nil {
		return nil, err
	}
	var (
		wg                sync.WaitGroup
		mu                sync.Mutex
		updatedWorkspaces []WorkspaceModel
		errChan           = make(chan error, len(workspaces))
	)

	processWorkspace := func(workspace WorkspaceModel) {
		defer wg.Done()
		updatedWorkspace, err := processWorkspaceActions(workspace, updateActionCompleted, accessToken)
		if err != nil {
			errChan <- err
			return
		}
		filter := bson.M{"_id": workspace.Id}
		update := bson.M{
			"$set": bson.M{"nodes": updatedWorkspace.Nodes},
		}
		updateResult, err := m.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			errChan <- err
			return
		}
		if updateResult.MatchedCount == 0 {
			errChan <- errWorkspaceNotFound
			return
		}
		mu.Lock()
		updatedWorkspaces = append(updatedWorkspaces, updatedWorkspace)
		mu.Unlock()
	}

	for _, workspace := range workspaces {
		wg.Add(1)
		go processWorkspace(workspace)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return updatedWorkspaces, nil
}

// isNodeCompleted checks if a given node is completed based on its status and action ID.
// It returns true if the node's status is "active" and its action ID matches the actionCompleted's action ID.
//
// Parameters:
//   - node: An ActionNodeModel representing the node to be checked.
//   - actionCompleted: An ActionCompletedModel containing the completed action's details. parent nodes are
//   - bool: true if the node is completed, false otherwise.
func isNodeCompleted(node ActionNodeModel, actionCompleted ActionCompletedModel) bool {
	return node.Status == "active" && node.ActionId == actionCompleted.ActionId
}

// isChildNodeReady checks if all parent nodes of a given node are in a "completed" status within a workspace.
// It iterates through the parents of the node and verifies their status in the workspace nodes.
// If any parent node is not in the "completed" status, it returns false. Otherwise, it returns true.
//
// Parameters:
// - node: The ActionNodeModel representing the node whose parents' statuses are being checked.
// - workspace: The WorkspaceModel containing the nodes and their statuses.
//
// Returns:
// - bool: True if all parent nodes are in the "completed" status, false otherwise.
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

// processWorkspaceActions processes the actions for a given workspace based on the completed action and updates the workspace status accordingly.
// It iterates through the nodes of the workspace, marking nodes as completed if they match the completed actions and setting their children to pending if they are ready.
// Finally, it initializes the workspace actions and returns the updated workspace.
//
// Parameters:
//   - workspace: The WorkspaceModel containing the nodes to be processed.
//   - actionCompleted: The ActionCompletedModel containing the action that has been completed.
//   - accessToken: A string representing the access token for initializing workspace actions.
//
// Returns:
//   - WorkspaceModel: The updated workspace model after processing the actions.
//   - error: An error if any occurred during the processing of the workspace actions.
func processWorkspaceActions(workspace WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) (WorkspaceModel, error) {
	for i, node := range workspace.Nodes {
		if isNodeCompleted(node, actionCompleted) {
			workspace.Nodes[i].Status = "completed"
			for j := range node.Children {
				if isChildNodeReady(workspace.Nodes[j], workspace) {
					err := initAction(workspace.Nodes[j], accessToken)
					if err == nil {
						log.Printf("Starting Action with id: %s", workspace.Nodes[j].ActionId.Hex())
						workspace.Nodes[j].Status = "active"
					} else {
						log.Printf("Error while starting Action with id: %s", workspace.Nodes[j].ActionId.Hex())
					}
				}
			}
		}
	}

	return workspace, nil
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
