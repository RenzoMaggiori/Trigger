package trigger

import (
	"context"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	token, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	event, ok := ctx.Value(SpotifyEventCtxKey).(ActionBody)
	if !ok {
		return errors.ErrEventCtx
	}

	action, _, err := action.GetByActionNameRequest(token, event.Type)
	if err != nil {
		return err
	}

	switch action.Action {
	case "watch_followers":
		data, ok := event.Data.(map[string]interface{})
		if !ok {
			return errors.ErrBadWebhookData
		}
		var followers FollowerChange
		if err := mapstructure.Decode(data, &followers); err != nil {
			return err
		}

		_, err := workspace.ActionCompletedRequest(token, workspace.ActionCompletedModel{
			ActionId: action.Id,
			Output: map[string]string{
				"followerCount": strconv.Itoa(followers.Followers),
				"increased":     strconv.FormatBool(followers.Increased),
			},
		})
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
