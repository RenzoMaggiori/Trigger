package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/spotify"
	"trigger.com/trigger/internal/sync"
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

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "spotify")
	if err != nil {
		return err
	}

	body, err := json.Marshal(struct{}{})
	if err != nil {
		return err
	}

	refreshToken := ""
	if syncModel.RefreshToken != nil {
		refreshToken = *syncModel.RefreshToken
	}
	client := spotify.Config().Client(ctx, &oauth2.Token{
		AccessToken:  syncModel.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		Expiry:       syncModel.Expiry,
		ExpiresIn:    syncModel.Expiry.Unix() - time.Now().Unix(),
	})

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPut,
			fmt.Sprintf("%s/me/player/play", spotify.BaseUrl),
			bytes.NewReader(body),
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.ErrSpotifyBadStatus
		}
		return fmt.Errorf("%w: %s", errors.ErrSpotifyBadStatus, body)
	}
	return nil
}
