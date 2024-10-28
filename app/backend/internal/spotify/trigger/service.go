package trigger

import (
	"context"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/errors"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	_, ok := ctx.Value(SpotifyEventCtxKey).(ActionBody)
	if !ok {
		return errors.ErrEventCtx
	}

	// TODO: handle webhook

	return nil
}

func (m Model) Stop(ctx context.Context) error {
	return nil
}
