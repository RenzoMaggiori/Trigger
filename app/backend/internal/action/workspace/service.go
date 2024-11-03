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
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
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
	if update.Name != nil {
		_, err := m.Collection.UpdateOne(ctx, bson.M{"_id": workspaceId}, bson.M{
			"$set": bson.M{
				"name": *update.Name,
			},
		})
		if err != nil {
			return nil, err
		}
	}

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
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)

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
		Name:   add.Name,
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

	return &newWorkspace, nil
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

func (m Model) initAction(ctx context.Context, workspace WorkspaceModel, actionNode ActionNodeModel, accessToken string) error {

	action, _, err := action.GetByIdRequest(accessToken, actionNode.ActionId.Hex())

	if err != nil {
		return err
	}

	actionCompleted, status, err := StartActionRequest(accessToken, actionNode, *action)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.ErrSettingAction
	}

	err = m.updateNodesStatus(workspace.UserId, actionNode.ActionId, "active")

	if err != nil {
		return err
	}

	if action.Type == "reaction" {
		if actionCompleted != nil {
			m.ActionCompleted(ctx, *actionCompleted)
		} else {
			m.ActionCompleted(ctx, ActionCompletedModel{
				ActionId:    action.Id,
				WorkspaceId: &workspace.Id,
				NodeId:      &actionNode.NodeId,
				Output:      make(map[string]string),
			})
		}
	}

	return nil
}

func (m Model) initWorkspace(ctx context.Context, workspace *WorkspaceModel, accessToken string, isNodeReady fnIsNodeReady) error {
	for _, node := range workspace.Nodes {
		if isNodeReady(node) {
			assignInputToAction(&node, workspace.Nodes)
			err := m.initAction(ctx, *workspace, node, accessToken)
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

func (m Model) WatchCompleted(ctx context.Context, watchCompleted WatchCompletedModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	user, _, err := user.GetUserByIdRequest(accessToken, session.UserId.Hex())

	if err != nil {
		return err
	}

	filter := bson.M{
		"user_id": user.Id,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": watchCompleted.ActionId,
				"node_id":   watchCompleted.NodeId,
			},
		},
	}

	update := bson.M{"$set": bson.M{}}

	for k, v := range watchCompleted.Input {
		update["$set"].(bson.M)[fmt.Sprintf("nodes.$.input.%s", k)] = v
	}

	res, err := m.Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.ErrWorkspaceNotFound
	}

	return nil
}

func isNodeReadyToStart(node ActionNodeModel, workspace WorkspaceModel) bool {
	if node.Status != "inactive" {
		return false
	}

	for _, parent := range node.Parents {
		for _, workspaceNode := range workspace.Nodes {
			if workspaceNode.NodeId == parent && workspaceNode.Status != "completed" {
				return false
			}
		}
	}
	return true
}

func (m Model) restartWorkspaces(ctx context.Context, workspaces []WorkspaceModel) {
	isWorkspaceCompleted := true

	for _, w := range workspaces {
		for _, node := range w.Nodes {
			if node.Status != "completed" {
				isWorkspaceCompleted = false
				break
			}
		}
		if isWorkspaceCompleted {
			log.Printf("Restarting workspace: %s", w.Id.Hex())
			_, err := m.Start(ctx, w.Id)
			if err != nil {
				log.Printf("Error restarting workspace: %s", w.Id.Hex())
			}
		} else {
			isWorkspaceCompleted = true
		}
	}
}

func (m Model) processWorkspace(ctx context.Context, workspace WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) error {

	filter := bson.M{
		"_id": workspace.Id,
		"nodes": bson.M{
			"$elemMatch": bson.M{
				"action_id": actionCompleted.ActionId,
			},
		},
	}

	if actionCompleted.NodeId != nil {
		filter["nodes"].(bson.M)["$elemMatch"].(bson.M)["node_id"] = actionCompleted.NodeId
	}

	update := bson.M{
		"$set": bson.M{
			"nodes.$.status": "completed",
		},
	}

	for key, value := range actionCompleted.Output {
		update["$set"].(bson.M)[fmt.Sprintf("nodes.$.output.%s", key)] = value
	}

	result, err := m.Collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
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
		ctx,
		updatedResult,
		accessToken,
		func(node ActionNodeModel) bool {
			return isNodeReadyToStart(node, *updatedResult)
		},
	)

	return nil
}

func (m Model) processUserWorkspaces(ctx context.Context, workspaces []WorkspaceModel, actionCompleted ActionCompletedModel, accessToken string) error {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, len(workspaces))
	)

	for _, workspace := range workspaces {
		wg.Add(1)
		go func(ws WorkspaceModel) {
			defer wg.Done()
			if err := m.processWorkspace(ctx, ws, actionCompleted, accessToken); err != nil {
				errChan <- err
			}
		}(workspace)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Model) ActionCompleted(ctx context.Context, actionCompleted ActionCompletedModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}
	user, _, err := user.GetUserByIdRequest(accessToken, session.UserId.Hex())

	if err != nil {
		return err
	}

	getWorkspaces := func() ([]WorkspaceModel, error) {
		userWorkspaces := make([]WorkspaceModel, 0)

		if actionCompleted.WorkspaceId != nil {
			w, err := m.GetById(ctx, *actionCompleted.WorkspaceId)
			if err != nil {
				return nil, err
			}
			userWorkspaces = append(userWorkspaces, *w)
		} else {
			userWorkspaces, err = m.GetByUserId(ctx, user.Id)
			if err != nil {
				return nil, err
			}
		}
		return userWorkspaces, nil
	}

	userWorkspaces, err := getWorkspaces()
	// Iterate over all user workspaces and update them in case they have any actions that are completed
	err = m.processUserWorkspaces(ctx, userWorkspaces, actionCompleted, accessToken)

	if err != nil {
		return nil
	}

	userWorkspaces, err = getWorkspaces()

	m.restartWorkspaces(ctx, userWorkspaces)

	return nil
}

func (m Model) Start(ctx context.Context, id primitive.ObjectID) (*WorkspaceModel, error) {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return nil, errors.ErrAccessTokenCtx
	}

	workspace, err := m.Stop(ctx, id)

	if err != nil {
		return nil, err
	}

	err = m.initWorkspace(ctx, workspace, accessToken, func(node ActionNodeModel) bool { return len(node.Parents) == 0 })

	if err != nil {
		return nil, err
	}
	err = m.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&workspace)
	if err != nil {
		return nil, err
	}

	return workspace, err
}

func (m Model) Stop(ctx context.Context, id primitive.ObjectID) (*WorkspaceModel, error) {
	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"nodes.$[].status": "inactive",
		},
	}

	res, err := m.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, errors.ErrWorkspaceNotFound
	}

	var workspace WorkspaceModel

	err = m.Collection.FindOne(ctx, filter).Decode(&workspace)

	if err != nil {
		return nil, err
	}

	//	TODO: call action stop functions
	return &workspace, nil
}

func (m Model) DeleteById(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}

	res, err := m.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.ErrWorkspaceNotFound
	}

	//	TODO: call action stop functions
	return nil
}

func (m Model) Templates(ctx context.Context) ([]AddWorkspaceModel, error) {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return nil, errors.ErrAccessTokenCtx
	}

	gmailWatch, _, err := action.GetByActionNameRequest(accessToken, "watch")
	if err != nil {
		return nil, err
	}

	gmailSend, _, err := action.GetByActionNameRequest(accessToken, "send_email")
	if err != nil {
		return nil, err
	}

	githubWatchPush, _, err := action.GetByActionNameRequest(accessToken, "watch_push")
	if err != nil {
		return nil, err
	}

	githubCreateIssue, _, err := action.GetByActionNameRequest(accessToken, "create_issue")
	if err != nil {
		return nil, err
	}

	spotifyWatchFollowers, _, err := action.GetByActionNameRequest(accessToken, "watch_followers")
	if err != nil {
		return nil, err
	}

	spotifyPlayMusic, _, err := action.GetByActionNameRequest(accessToken, "play_music")
	if err != nil {
		return nil, err
	}

	twitchChannelFollow, _, err := action.GetByActionNameRequest(accessToken, "channel_follow")
	if err != nil {
		return nil, err
	}

	twitchSendChatMessage, _, err := action.GetByActionNameRequest(accessToken, "send_chat_message")
	if err != nil {
		return nil, err
	}

	discordWatchChannelMessage, _, err := action.GetByActionNameRequest(accessToken, "watch_channel_message")
	if err != nil {
		return nil, err
	}

	discordSendChannelMessage, _, err := action.GetByActionNameRequest(accessToken, "send_channel_message")
	if err != nil {
		return nil, err
	}

	bitbucketWatchIssueCreated, _, err := action.GetByActionNameRequest(accessToken, "watch_issue_created")
	if err != nil {
		return nil, err
	}

	bitbucketCreatePullRequest, _, err := action.GetByActionNameRequest(accessToken, "create_pull_request")
	if err != nil {
		return nil, err
	}

	templates := []AddWorkspaceModel{
		{
			Name: "gmail",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Gmail-0",
					ActionId: gmailWatch.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Gmail-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Gmail-1",
					ActionId: gmailSend.Id,
					Input:    map[string]string{},
					Parents:  []string{"Gmail-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
		{
			Name: "github",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Github-0",
					ActionId: githubWatchPush.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Github-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Github-1",
					ActionId: githubCreateIssue.Id,
					Input:    map[string]string{},
					Parents:  []string{"Github-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
		{
			Name: "spotify",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Spotify-0",
					ActionId: spotifyWatchFollowers.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Spotify-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Spotify-1",
					ActionId: spotifyPlayMusic.Id,
					Input:    map[string]string{},
					Parents:  []string{"Spotify-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
		{
			Name: "twitch",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Twitch-0",
					ActionId: twitchChannelFollow.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Twitch-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Twitch-1",
					ActionId: twitchSendChatMessage.Id,
					Input:    map[string]string{},
					Parents:  []string{"Twitch-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
		{
			Name: "discord",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Discord-0",
					ActionId: discordWatchChannelMessage.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Discord-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Discord-1",
					ActionId: discordSendChannelMessage.Id,
					Input:    map[string]string{},
					Parents:  []string{"Discord-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
		{
			Name: "bitbucket",
			Nodes: []AddActionNodeModel{
				{
					NodeId:   "Bitbucket-0",
					ActionId: bitbucketWatchIssueCreated.Id,
					Input:    map[string]string{},
					Parents:  make([]string, 0),
					Children: []string{"Bitbucket-1"},
					XPos:     100,
					YPos:     100,
				},
				{
					NodeId:   "Bitbucket-1",
					ActionId: bitbucketCreatePullRequest.Id,
					Input:    map[string]string{},
					Parents:  []string{"Bitbucket-0"},
					Children: make([]string, 0),
					XPos:     100,
					YPos:     200,
				},
			},
		},
	}
	return templates, nil
}
