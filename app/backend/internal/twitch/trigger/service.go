package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/twitch"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	userResponse, err := twitch.GetUserByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	appAccessToken, err := twitch.GetAppAccessTokenrRequest()
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
			Callback: fmt.Sprintf("%s/api/twitch/trigger/webhook", os.Getenv("TWITCH_BASE_URL")),
			Secret:   os.Getenv("TWITCH_CLIENT_SECRET"),
		},
	}
	body, err := json.Marshal(watchBody)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
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
	webhookVerfication, ok := ctx.Value(WebhookVerificationCtxKey).(WebhookVerificationRequest)
	if !ok {
		return errors.ErrWebhookVerificationCtx
	}

	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
