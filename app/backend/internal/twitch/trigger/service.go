package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/internal/twitch"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	triggerUser, _, err := user.GetUserByAccesstokenRequest(accessToken)
	if err != nil {
		return err
	}

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, triggerUser.Id.Hex(), "twitch")
	if err != nil {
		return err
	}

	userResponse, err := twitch.GetUserByAccessTokenRequest(*syncModel)
	if err != nil {
		return err
	}

	appAccessToken, err := twitch.GetAppAccessTokenRequest()
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
			Callback: fmt.Sprintf("%s/api/twitch/trigger/webhook?userId=%s", os.Getenv("SERVER_BASE_URL"), triggerUser.Id.Hex()),
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
		return errors.ErrTwitchWatch
	}

	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	userId, ok := ctx.Value(WebhookUserIdCtxKey).(string)
	if !ok {
		return errors.ErrBadUserId
	}

	// webhookVerfication, ok := ctx.Value(WebhookVerificationCtxKey).(WebhookVerificationRequest)
	// if !ok {
	// 	return errors.ErrWebhookVerificationCtx
	// }

	userSesssion, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), userId)
	if err != nil {
		return err
	}
	if len(userSesssion) == 0 {
		return errors.ErrSessionNotFound
	}

	accessToken := userSesssion[0].AccessToken
	action, _, err := action.GetByActionNameRequest(accessToken, "watch_channel_follow")
	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		Output:   map[string]string{},
	}
	_, err = workspace.ActionCompletedRequest(accessToken, update)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
