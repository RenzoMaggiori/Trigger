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

type Reaction interface {
	Reaction(ctx context.Context, action workspace.ActionNodeModel) error
}

type MultipleReactions interface {
	MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error
}
