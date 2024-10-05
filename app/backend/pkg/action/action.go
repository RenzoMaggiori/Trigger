package action

import (
	"context"

	"trigger.com/trigger/internal/action/workspace"
)

type Action interface {
	Watch(ctx context.Context, action workspace.ActionNodeModel) error
	Webhook(ctx context.Context) error
}
