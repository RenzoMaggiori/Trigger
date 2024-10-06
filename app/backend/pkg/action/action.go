package action

import (
	"context"

	"trigger.com/trigger/internal/action/workspace"
)

type Trigger interface {
	Watch(ctx context.Context, action workspace.ActionNodeModel) error
	Webhook(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Action interface {
	Action(ctx context.Context, action workspace.ActionNodeModel) error
}
