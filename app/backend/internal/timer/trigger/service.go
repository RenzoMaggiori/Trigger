package trigger

import (
	"context"
	"time"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/user"
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

	event, ok := ctx.Value(timerActionCtxKey).(ActionBody)
	if !ok {
		return errors.ErrEventCtx
	}

	user, _, err := user.GetUserByAccesstokenRequest(token)
	if err != nil {
		return err
	}

	action, _, err := action.GetByActionNameRequest(token, event.Name)
	if err != nil {
		return err
	}

	_, err = workspace.ActionCompletedRequest(token, workspace.ActionCompletedModel{
		UserId:   user.Id,
		ActionId: action.Id,
		Output: map[string]string{
			"datetime": event.DateTime.Format(time.RFC850),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
