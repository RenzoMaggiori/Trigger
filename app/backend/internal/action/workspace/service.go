package workspace

import (
	"context"
	"fmt"
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

func (m Model) updateNodesStatus(actionId string, userId string, status string) error {
	filter := bson.M{
		"user_id": userId,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": actionId,
			},
		},
	}

	// Define the update: set the output field for the matching nodes
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
			m.updateNodesStatus(node.ActionId.Hex(), workspace.UserId.Hex(), "active")
		}
	}
	return nil
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
			Output:   map[string]string{},
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

	// Construct the update for each key in the Output map
	outputUpdate := bson.M{}
	for key, value := range actionCompleted.Output {
		outputUpdate[fmt.Sprintf("nodes.$.output.%s", key)] = value
	}

	// Use $set to apply the updated Output fields without overwriting the entire map
	update := bson.M{
		"$set": bson.M{
			"nodes.$.status": "completed",
		},
	}
	// Merge the outputUpdate into the main $set update
	for k, v := range outputUpdate {
		update["$set"].(bson.M)[k] = v
	}

	result, err := m.Collection.UpdateMany(context.TODO(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return errors.ErrUpdatingWorkspace
	}

	updatedResult, err := m.GetById(context.TODO(), workspace.Id)

	if err != nil {
		return err
	}

	m.initWorkspace(
		updatedResult,
		accessToken,
		func(node ActionNodeModel) bool {
			return isNodeReady(node, *updatedResult)
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

// Check if the node is ready to be started.
//
// In this case we are checking if all parent nodes are completed (this logic can be changed)
func isNodeReady(child ActionNodeModel, workspace WorkspaceModel) bool {
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

func (m Model) UpdateById(ctx context.Context, id primitive.ObjectID, update *UpdateWorkspaceModel) (*WorkspaceModel, error) {
	// Create filter to find document by ID
	filter := bson.M{"_id": id}

	// Iterate over each node to build individual update operations
	for _, node := range update.Nodes {
		// Filter to match specific node within nodes array by node_id
		nodeFilter := bson.M{
			"_id":           id,
			"nodes.node_id": node.NodeId,
		}

		// Build update for the specific node fields
		nodeUpdate := bson.M{
			"$set": bson.M{
				"nodes.$.input":    node.Input,
				"nodes.$.output":   node.Output,
				"nodes.$.parents":  node.Parents,
				"nodes.$.children": node.Children,
				"nodes.$.x_pos":    node.XPos,
				"nodes.$.y_pos":    node.YPos,
				"nodes.$.status":   "completed", // Example: setting a status on each node
			},
		}

		// Execute update for the current node
		_, err := m.Collection.UpdateOne(ctx, nodeFilter, nodeUpdate)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", errors.ErrWorkspaceNotFound, err)
		}
	}

	// Retrieve the updated document
	var updatedWorkspace WorkspaceModel
	err := m.Collection.FindOne(ctx, filter).Decode(&updatedWorkspace)
	if err != nil {
		return nil, err
	}

	return &updatedWorkspace, nil
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

	workspaces := make([]WorkspaceModel, 0)
	if err = cursor.All(ctx, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}
