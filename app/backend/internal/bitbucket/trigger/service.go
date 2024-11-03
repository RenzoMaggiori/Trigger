package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/bitbucket"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "bitbucket")
	if err != nil {
		return err
	}

	client, err := oaclient.New(ctx, bitbucket.Config(), syncModel)
	if err != nil {
		return err
	}

	workspace, ok := actionNode.Input["workspace"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}
	repository, ok := actionNode.Input["repository"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	watchBody := WatchBody{
		Description: "This is a watch",
		URL: fmt.Sprintf("%s/api/bitbucket/trigger/webhook?user_id=%s",
			os.Getenv("SERVER_BASE_URL"),
			session.UserId.Hex()),
		Active: true,
		Secret: os.Getenv("BITBUCKET_SECRET"),
		Events: []string{"issue:created", "repo:push"},
	}

	body, err := json.Marshal(watchBody)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/hooks", workspace, repository),
			bytes.NewReader(body),
			nil,
		),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return errors.ErrBitbucketBadStatus
	}

	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	userId, ok := ctx.Value(userIdCtxKey).(string)
	if !ok {
		return errors.ErrBadUserId
	}

	webhookRequest, ok := ctx.Value(WebhookEventCtxKey).(WebhookRequest)

	if !ok {
		return errors.ErrGithubCommitData
	}

	session, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return err
	}
	if len(session) == 0 {
		return errors.ErrSessionNotFound
	}

	action, _, err := action.GetByActionNameRequest(session[0].AccessToken, "watch_issue_created")
	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		Output: map[string]string{
			"title":   webhookRequest.Issue.Title,
			"content": webhookRequest.Issue.Content.Raw,
		},
	}
	_, err = workspace.ActionCompletedRequest(session[0].AccessToken, update)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
