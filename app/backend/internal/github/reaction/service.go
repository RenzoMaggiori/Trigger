package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	githubClient "github.com/google/go-github/v66/github"

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

func (m Model) Reaction(ctx context.Context, actionNode workspace.ActionNodeModel) error {
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

	githubClient := githubClient.NewClient(nil).WithAuthToken(syncModel.AccessToken)
	githubUser, _, err := githubClient.Users.Get(ctx, "")

	if err != nil {
		return err
	}

	owner := *githubUser.Login
	repo := actionNode.Input["repo"]
	body, err := json.Marshal(map[string]any{
		"title":     actionNode.Input["title"],
		"body":      actionNode.Input["body"],
		"assignees": []string{owner},
		"labels":    []string{actionNode.Input["labels"]},
	})
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/repos/%s/%s/issues", githuBaseUrl, owner, repo),
			bytes.NewReader(body),
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
