package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"

	"trigger.com/trigger/pkg/decode"
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
		return errors.ErrAccessTokenCtxKey
	}

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/session/access_token/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), accessToken),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.ErrSessionNotFound
	}

	session, err := decode.Json[session.SessionModel](res.Body)
	if err != nil {
		return err
	}

	res, err = fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/%s", os.Getenv("USER_SERVICE_BASE_URL"), session.UserId),
			nil,
			map[string]string{
				"Authorization": accessToken,
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.ErrUserNotFound
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		return err
	}
	// TODO: get correct access token from sync service

	if len(actionNode.Output) != 2 {
		return errors.ErrInvalidReactionOuput
	}

	body, err := json.Marshal(map[string]any{
		"title":     "Reaction Title",
		"body":      "Reaction Body",
		"assignees": []string{user.Email},
		"milestone": 1,
		"labels":    []string{"bug"},
	})
	if err != nil {
		return err
	}

	owner := actionNode.Input[0]
	repo := actionNode.Input[1]
	res, err = fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/repos/%s/%s/issues", githuBaseUrl, owner, repo),
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
