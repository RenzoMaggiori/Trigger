package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/github"
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
		return errors.ErrAccessTokenCtxKey
	}

	body, err := json.Marshal(map[string]any{
		"name":   "web",
		"active": true,
		"events": []string{"push"},
		"config": map[string]any{
			"url":          os.Getenv("GITHUB_WEBHOOK_URL"),
			"content_type": "json",
			"insecure_ssl": "0",
		},
	})
	if err != nil {
		return err
	}

	if len(actionNode.Input) != 2 {
		return errors.ErrInvalidReactionInput
	}

	// TODO: get correct access token from sync service
	owner := actionNode.Input[0]
	repo := actionNode.Input[1]
	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/repos/%s/%s/hooks", githuBaseUrl, owner, repo),
			bytes.NewReader(body),
			map[string]string{
				"Authorization":        fmt.Sprintf("Bearer %s", accessToken),
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

func (m Model) Stop(ctx context.Context) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtxKey
	}

	body, ok := ctx.Value(github.StopCtxKey).(StopModel)
	if !ok {
		return errors.ErrGithubStopModelNotFound
	}

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/repos/%s/%s/hooks/%s", githuBaseUrl, body.Owner, body.Repo, body.HookId),
			nil,
			map[string]string{
				"Authorization":        fmt.Sprintf("Bearer %s", accessToken),
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

func (m Model) Webhook(ctx context.Context) error {
	/* update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		UserId:   user.Id,
		Output:   map[string]any{"hello": "world"},
	}

	_, err = workspace.ActionCompletedRequest(googleSession.AccessToken, update)

	if err != nil {
		return err
	}

	return nil */
	return nil
}
