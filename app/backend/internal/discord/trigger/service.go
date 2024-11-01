package trigger

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"

	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	// accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	// if !ok {
	// 	return errors.ErrAccessTokenCtx
	// }

	// session, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	// if err != nil {
	// 	return err
	// }

	// userId := session.UserId

	// channel_id := actionNode.Input["channel_id"]

	// actionId := actionNode.ActionId.Hex()

	// discord_me, err := m.GetMe(accessToken)
	// if err != nil {
	// 	return err
	// }

	// existingSession, _ := m.GetSessionByUserId(session.UserId)
	// if existingSession != nil {
	// 	if existingSession.ChannelId == channel_id {
	// 		log.Println("Session already exists for this channel")
	// 		return nil
	// 	}
	// }

	// newSession := &DiscordSessionModel{
	// 	UserId:    userId,
	// 	ChannelId: channel_id,
	// 	ActionId:  actionId,
	// 	Token:     session.AccessToken,
	// 	DiscordData: discord_me,
	// }
	// err = worker.AddSession(newSession)
	// if err != nil {
	// 	log.Printf("Error adding session [%s]: %v", newSession.ChannelId, err)
	// 	return err
	// }
    return nil
}

func (m *Model) Webhook(ctx context.Context) error {
	token, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	event, ok := ctx.Value(DiscordEventCtxKey).(ActionBody)
	if !ok {
		return errors.ErrEventCtx
	}

	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return err
	}

	action, _, err := action.GetByActionNameRequest(token, event.Type)
	if err != nil {
		return err
	}

	switch action.Action {
	case "watch_channel_message":
		data, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.ErrBadWebhookData
		}
		var msgInfo MsgInfo
		if err := mapstructure.Decode(data, &msgInfo); err != nil {
			return err
		}

		_, err := workspace.ActionCompletedRequest(token, workspace.ActionCompletedModel{
			UserId:   user.Id,
			ActionId: action.Id,
			Output: map[string]string{
                "content": msgInfo.Content,
				"author_id": msgInfo.AuthoId,
				"author_username":  msgInfo.AuthoUsername,
            },
			// NodeId:   actionNode.Id,
		})
		return err
	}
	return nil
}

func (m *Model) Stop(ctx context.Context) error {
    return nil
}
