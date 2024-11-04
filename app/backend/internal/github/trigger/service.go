package trigger

import (
	"context"
	"fmt"
	"os"

	githubClient "github.com/google/go-github/v66/github"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/github"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

const (
	githuBaseUrl string = "https://api.github.com"
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

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "github")
	if err != nil {
		return err
	}

	client, err := oaclient.New(ctx, github.Config(), syncModel)
	if err != nil {
		return err
	}

	ghClient := githubClient.NewClient(client)
	owner, ok := actionNode.Input["owner"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	repo, ok := actionNode.Input["repo"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	name := "web"
	active := true
	url := fmt.Sprintf("%s/api/github/trigger/webhook?userId=%s", os.Getenv("SERVER_BASE_URL"), syncModel.UserId.Hex())
	contentType := "json"
	insecureSSL := "0"
	_, githubRes, err := ghClient.Repositories.CreateHook(
		ctx,
		owner,
		repo,
		&githubClient.Hook{
			Name:   &name,
			Active: &active,
			Events: []string{"push"},
			Config: &githubClient.HookConfig{
				URL:         &url,
				ContentType: &contentType,
				InsecureSSL: &insecureSSL,
			},
		},
	)

	if githubRes != nil && githubRes.StatusCode == 422 {
		return nil
	}

	if err != nil {
		return err
	}
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	userId, ok := ctx.Value(userIdCtxKey).(string)
	if !ok {
		return errors.ErrBadUserId
	}

	commit, ok := ctx.Value(GithubCommitCtxKey).(githubClient.PushEvent)
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

	action, _, err := action.GetByActionNameRequest(session[0].AccessToken, "watch_push")
	if err != nil {
		return err
	}

	author := ""
	message := ""
	if len(commit.Commits) != 0 {
		if commit.Commits[0].Author != nil && commit.Commits[0].Author.Name != nil {
			author = *commit.Commits[0].Author.Name
		}
		if commit.Commits[0].Message != nil {
			message = *commit.Commits[0].Message
		}
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		Output: map[string]string{
			"author":  author,
			"message": message,
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
