package reaction

import (
	"context"

	githubClient "github.com/google/go-github/v66/github"

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

	ghClient := githubClient.NewClient(client)
	githubUser, _, err := ghClient.Users.Get(ctx, "")
	if err != nil {
		return err
	}

	owner := *githubUser.Login
	repo, ok := actionNode.Input["repo"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	title, ok := actionNode.Input["title"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	body, ok := actionNode.Input["description"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	labels := []string{"bug"}
	_, _, err = ghClient.Issues.Create(ctx, owner, repo, &githubClient.IssueRequest{
		Title:  &title,
		Body:   &body,
		Labels: &labels,
	})
	if err != nil {
		return err
	}
	return nil
}
