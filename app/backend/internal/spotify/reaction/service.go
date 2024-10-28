package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"

	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	switch actionName {
	case "play_music":
		return m.PlayMusic(ctx, accessToken, action)
	}

	return nil
}

func (m Model) PlayMusic(ctx context.Context, accessToken string, actionNode workspace.ActionNodeModel) error {
	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	user, _, err := user.GetUserByIdRequest(accessToken, session.UserId.Hex())

	if err != nil {
		return err
	}

	// TODO: play music

	return nil
}
