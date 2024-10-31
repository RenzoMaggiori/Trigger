package trigger

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"

	// "trigger.com/trigger/internal/discord/worker"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (m *Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
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
	case "watch_message":
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
				"author":  msgInfo.Author,
                "content": msgInfo.Content,
            },
		})
		return err
	}
	return nil
}

func (m *Model) Stop(ctx context.Context) error {
    return nil
}
