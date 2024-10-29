package workspace

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceCtx string

const WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")

const AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")

type Service interface {
	Get(context.Context) ([]WorkspaceModel, error)
	GetById(context.Context, primitive.ObjectID) (*WorkspaceModel, error)
	GetByUserId(context.Context, primitive.ObjectID) ([]WorkspaceModel, error)
	GetByActionId(context.Context, primitive.ObjectID) ([]WorkspaceModel, error)
	Add(context.Context, *AddWorkspaceModel) (*WorkspaceModel, error)
	ActionCompleted(context.Context, ActionCompletedModel) ([]WorkspaceModel, error)
	UpdateById(context.Context, primitive.ObjectID, *UpdateWorkspaceModel) (*WorkspaceModel, error)
	WatchCompleted(context.Context, WatchCompletedModel) ([]WorkspaceModel, error)
	// DeleteById(context.Context, primitive.ObjectID) error
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type WorkspaceModel struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Nodes  []ActionNodeModel  `json:"nodes" bson:"nodes"`
}

// status : completed / active / pending / inactive
// completed: it is what it says on the tin
// active: waiting for an action to happen
// pending: waiting for activating the action
// inactive: depends on other actions / triggers

type ActionNodeModel struct {
	NodeId   string             `json:"node_id" bson:"node_id"`
	ActionId primitive.ObjectID `json:"action_id" bson:"action_id"`
	Input    map[string]string  `json:"input" bson:"input"`
	Output   map[string]string  `json:"output" bson:"output"`
	Parents  []string           `json:"parents" bson:"parents"`
	Children []string           `json:"children" bson:"children"`
	Status   string             `json:"status" bson:"status"`
	XPos     float32            `json:"x_pos" bson:"x_pos"`
	YPos     float32            `json:"y_pos" bson:"y_pos"`
}

type AddWorkspaceModel struct {
	Nodes []AddActionNodeModel `json:"nodes" bson:"nodes"`
}

type AddActionNodeModel struct {
	NodeId   string             `json:"node_id" bson:"node_id"`
	ActionId primitive.ObjectID `json:"action_id" bson:"action_id"`
	Input    map[string]string  `json:"input" bson:"input"`
	Parents  []string           `json:"parents" bson:"parents"`
	Children []string           `json:"children" bson:"children"`
	XPos     float32            `json:"x_pos" bson:"x_pos"`
	YPos     float32            `json:"y_pos" bson:"y_pos"`
}

type UpdateActionNodeModel struct {
	NodeId   string            `json:"node_id" bson:"node_id"`
	Input    map[string]string `json:"fields" bson:"fields"`
	Output   map[string]string `json:"output" bson:"input"`
	Parents  []string          `json:"parents" bson:"parents"`
	Children []string          `json:"children" bson:"children"`
	XPos     float32           `json:"x_pos" bson:"x_pos"`
	YPos     float32           `json:"y_pos" bson:"y_pos"`
}

type UpdateWorkspaceModel struct {
	Nodes []UpdateActionNodeModel `json:"nodes" bson:"nodes"`
}

type ActionCompletedModel struct {
	ActionId primitive.ObjectID `json:"action_id"`
	UserId   primitive.ObjectID `json:"user_id"`
	Output   map[string]string  `json:"output"`
}

type WatchCompletedModel struct {
	UserId   primitive.ObjectID `json:"user_id"`
	ActionId primitive.ObjectID `json:"action_id"`
	Output   map[string]string  `json:"output"`
}
