package trigger

import (
	"context"
	"fmt"
	"net/http"
	"time"

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

	watchCompleted := workspace.WatchCompletedModel{
		ActionId: actionNode.ActionId,
		UserId:   session.UserId,
		Input: map[string]string{
			"since": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	_, _, err = workspace.WatchCompletedRequest(accessToken, watchCompleted)
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

	user, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	action, _, err := action.GetByActionNameRequest(accessToken, "watch_commit")

	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		UserId:   user.Id,
		Output:   map[string]string{"author": *commit.Commit.Author.Name},
	}

	_, err = workspace.ActionCompletedRequest(accessToken, update)

	if err != nil {
		return err
	}

	return nil
}

func (m Model) Stop(ctx context.Context) error {
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

	body, ok := ctx.Value(github.StopCtxKey).(StopModel)
	if !ok {
		return errors.ErrGithubStopModelNotFound
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/repos/%s/%s/hooks/%s", githuBaseUrl, body.Owner, body.Repo, body.HookId),
			nil,
			map[string]string{
				"Accept":               "application/vnd.github+json",
				"X-GitHub-Api-Version": "2022-11-28",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return fmt.Errorf("%w: received %s", errors.ErrInvalidGithubStatus, res.Status)
	}
	return nil
}
