package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/twitch"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error {
	switch actionName {
	case "send_channel_message":
		return m.SendChannelMessage(ctx, action)
	}

	return nil
}

func (m Model) SendChannelMessage(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)

	if !ok {
		return errors.ErrAccessTokenCtx
	}

	user, err := twitch.GetUserByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	sendChannelMessageBody := SendChannelMessageBody{
		BroadcasterId: user.Data[0].ID,
		SenderId:      user.Data[0].ID,
		Message:       actionNode.Input["message"],
	}

	body, err := json.Marshal(sendChannelMessageBody)

	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://api.twitch.tv/helix/chat/messages",
			bytes.NewReader(body),
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
	if res.StatusCode >= 400 {
		return errors.ErrTwitchSendMessage
	}

	messageSent, err := decode.Json[MessageData](res.Body)

	if err != nil {
		return err
	}

	if len(messageSent.Data) == 0 || !messageSent.Data[0].IsSent {
		return errors.ErrTwitchSendMessage
	}

	return nil

}
