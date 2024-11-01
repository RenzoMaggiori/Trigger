package trigger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	githubClient "github.com/google/go-github/v66/github"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/github"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
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
	url := ""
	contentType := "json"
	insecureSSL := "0"
	_, _, err = ghClient.Repositories.CreateHook(
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
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	commit, ok := ctx.Value(GithubCommitCtxKey).(githubClient.RepositoryCommit)
	if !ok {
		return errors.ErrGithubCommitData
	}

	sesion, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	action, _, err := action.GetByActionNameRequest(accessToken, "watch_commit")
	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		UserId:   sesion.UserId,
		Output: map[string]string{
			"author":  *commit.Commit.Author.Name,
			"message": *commit.Commit.Message,
		},
	}
	_, err = workspace.ActionCompletedRequest(accessToken, update)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
