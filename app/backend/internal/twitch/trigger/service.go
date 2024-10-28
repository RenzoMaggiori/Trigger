package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func getTwitchAppAccessToken() (*AppAccessTokenResponse, error) {
	appAccessTokenBody := AppAccessTokenBody{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		GrantType:    "client_credentials",
	}

	body, err := json.Marshal(appAccessTokenBody)

	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://id.twitch.tv/oauth2/token",
			bytes.NewReader(body),
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		log.Printf("Watch body: %s", bodyBytes)
		return nil, errors.ErrTwitchAppAccessToken
	}

	appAccessTokenResponse, err := decode.Json[AppAccessTokenResponse](res.Body)

	if err != nil {
		return nil, err
	}

	return &appAccessTokenResponse, nil
}

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)

	if !ok {
		return errors.ErrAccessTokenCtx
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("https://api.twitch.tv/helix/users?login=%s", actionNode.Input["channel"]),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				"Client-Id":     os.Getenv("TWITCH_CLIENT_ID"),
				"Content-Type":  "application/json",
			},
		),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.ErrTwitchUser
	}

	userResponse, err := decode.Json[UserResponse](res.Body)

	if err != nil {
		return err
	}

	if len(userResponse.Data) == 0 {
		return errors.ErrTwitchUser
	}

	appAccessToken, err := getTwitchAppAccessToken()

	if err != nil {
		return err
	}

	watchBody := ChannelFollowSubscriptionBody{
		Type:    "channel.follow",
		Version: "2",
		Condition: ChannelFollowCondition{
			BroadcasterUserID: userResponse.Data[0].ID,
			ModeratorUserID:   userResponse.Data[0].ID,
		},
		Transport: ChannelFollowTransport{
			Method:   "webhook",
			Callback: "https://7bea-163-5-23-104.ngrok-free.app/api/twitch/trigger/webhook",
			Secret:   os.Getenv("TWITCH_CLIENT_SECRET"),
		},
	}

	body, err := json.Marshal(watchBody)

	if err != nil {
		return err
	}

	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://api.twitch.tv/helix/eventsub/subscriptions",
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", appAccessToken.AccessToken),
				"Client-Id":     os.Getenv("TWITCH_CLIENT_ID"),
				"Content-Type":  "application/json",
			},
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
		log.Printf("Watch body: %s", bodyBytes)
		return errors.ErrGmailWatch
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	watchCompleted := workspace.WatchCompletedModel{
		ActionId: actionNode.ActionId,
		UserId:   session.UserId,
	}

	_, _, err = workspace.WatchCompletedRequest(accessToken, watchCompleted)

	if err != nil {
		return err
	}

	return nil
}

func (m Model) Webhook(ctx context.Context) error {

	return nil
}

func (m Model) Stop(ctx context.Context) error {

	return nil
}
