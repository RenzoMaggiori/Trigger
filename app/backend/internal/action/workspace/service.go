package workspace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
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

func initWorkspaceActions(workspace WorkspaceModel, accessToken string) error {

	for _, node := range workspace.Nodes {
		if node.Status == "active" {
			res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
				http.MethodGet,
				fmt.Sprintf("%s/api/action/id/%s", os.Getenv("ACTION_SERVICE_BASE_URL"), node.ActionId.Hex()),
				nil,
				map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				},
			))

			if err != nil {
				return errFetchingActions
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				return errFetchingActions
			}
			action, err := decode.Json[action.ActionModel](res.Body)

			if err != nil {
				return errActionTypeNone
			}
			actionEnv := fmt.Sprintf("%s_SERVICE_BASE_URL", strings.ToUpper(action.Provider))

			body, err := json.Marshal(node)

			if err != nil {
				return err
			}
			// Call the reaction / trigger
			res, err = fetch.Fetch(
				http.DefaultClient,
				fetch.NewFetchRequest(
					http.MethodPost,
					fmt.Sprintf("%s/api/services/%s/%s/%s", os.Getenv(actionEnv), action.Provider, action.Type, action.Action),
					bytes.NewReader(body),
					map[string]string{
						"Authorization": fmt.Sprintf("Bearer %s", accessToken),
					},
				),
			)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				return errAction
			}
		}
	}
	return nil
}

func nodeStatus(parents []string) string {
	if len(parents) == 0 {
		return "active"
	}
	return "inactive"
}

func (m Model) Add(ctx context.Context, add *AddWorkspaceModel) (*WorkspaceModel, error) {
	accessToken := ctx.Value(AccessTokenCtxKey).(string)
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/access_token/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), accessToken),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, errSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errSessionNotFound
	}
	session, err := decode.Json[session.SessionModel](res.Body)

	if err != nil {
		return nil, errSessionTypeNone
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
			Fields:   node.Fields,
			Parents:  node.Parents,
			Children: node.Children,
			Status:   nodeStatus(node.Parents),
			XPos:     node.XPos,
			YPos:     node.YPos,
		}
		newWorkspace.Nodes = append(newWorkspace.Nodes, node)
	}

	_, err = m.Collection.InsertOne(ctx, newWorkspace)
	if err != nil {
		return nil, errCreatingWorkspace
	}

	err = initWorkspaceActions(newWorkspace, accessToken)
	if err != nil {
		return nil, err
	}

	return &newWorkspace, nil
}

func (m Model) UpdateActionCompleted(ctx context.Context,
	id primitive.ObjectID, nodeId string) (*WorkspaceModel, error) {

	var workspace WorkspaceModel

	filter := bson.M{"_id": id}
	err := m.Collection.FindOne(ctx, filter).Decode(&workspace)
	if err != nil {
		return nil, errWorkspaceNotFound
	}

	for i, node := range workspace.Nodes {
		if node.NodeId == nodeId {
			workspace.Nodes[i].Status = "completed"
			for _, child := range node.Children {
				for j := range workspace.Nodes {
					if workspace.Nodes[j].NodeId == child {
						workspace.Nodes[j].Status = "active"
					}
				}
			}
			break
		}
	}
	update := bson.M{
		"$set": bson.M{
			"nodes": workspace.Nodes,
		},
	}
	updateResult, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if updateResult.MatchedCount == 0 {
		return nil, errWorkspaceNotFound
	}
	err = initWorkspaceActions(workspace, ctx.Value(AccessTokenCtxKey).(string))

	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

// func (m Model) UpdateById(ctx context.Context, id primitive.ObjectID, update *UpdateWorkspaceModel) (*WorkspaceModel, error) {
// 	filter := bson.M{"_id": id, "solved": false}
// 	updateData := bson.M{"$set": update}

// 	_, err := m.Collection.UpdateOne(ctx, filter, updateData)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %v", errUserActionNotFound, err)
// 	}

// 	var updatedUserAction WorkspaceModel
// 	err = m.Collection.FindOne(ctx, filter).Decode(&updatedUserAction)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &updatedUserAction, nil

// }

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
