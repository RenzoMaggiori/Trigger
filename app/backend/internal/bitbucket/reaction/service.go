package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/bitbucket"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"

	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
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
	title, ok := actionNode.Input["title"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}
	sourceBranch, ok := actionNode.Input["source_branch"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	destinationBranch, ok := actionNode.Input["destination_branch"]
	if !ok {
		return errors.ErrInvalidReactionInput
	}

	pullRequest := PullRequest{
		Title:       title,
		Source:      BranchInfo{Branch: Branch{Name: sourceBranch}},
		Destination: BranchInfo{Branch: Branch{Name: destinationBranch}},
	}

	body, err := json.Marshal(pullRequest)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/pullrequests", workspace, repository),
			bytes.NewReader(body),
			map[string]string{"Content-Type": "application/json"},
		),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v: %s", errors.ErrBitbucketBadStatus, bodyBytes)
	}

	return nil
}
