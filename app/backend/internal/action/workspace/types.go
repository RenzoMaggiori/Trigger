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
	Add(context.Context, *AddWorkspaceModel) (*WorkspaceModel, error)
	// UpdateById(context.Context, primitive.ObjectID, *UpdateUserActionModel) (*UserActionModel, error)
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
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Nodes  []ActionNodeModel  `json:"nodes" bson:"nodes"`
}

type ActionNodeModel struct {
	NodeId   string             `json:"node_id" bson:"node_id"`
	ActionId primitive.ObjectID `json:"action_id" bson:"action_id"`
	Fields   []any              `json:"fields" bson:"fields"`
	Parents  []string           `json:"parents" bson:"parents"`
	Children []string           `json:"children" bson:"children"`
	Solved   bool               `json:"solved" bson:"solved"`
	Active   bool               `json:"active" bson:"active"`
	XPos     float32            `json:"x_pos" bson:"x_pos"`
	YPos     float32            `json:"y_pos" bson:"y_pos"`
}

type AddActionNodeModel struct {
	NodeId   string             `json:"node_id" bson:"node_id"`
	ActionId primitive.ObjectID `json:"action_id" bson:"action_id"`
	Fields   []any              `json:"fields" bson:"fields"`
	Parents  []string           `json:"parents" bson:"parents"`
	Children []string           `json:"children" bson:"children"`
	XPos     float32            `json:"x_pos" bson:"x_pos"`
	YPos     float32            `json:"y_pos" bson:"y_pos"`
}

type AddWorkspaceModel struct {
	Nodes []AddActionNodeModel `json:"nodes" bson:"nodes"`
}

// type UpdateActionNode struct {
// 	ActionId   string   `json:"id" bson:"action_id"`
// 	ActionType string   `json:"action_type" bson:"action_type"`
// 	Fields     []any    `json:"fields" bson:"fields"`
// 	Parents    []string `json:"parents" bson:"parents"`
// 	Children   []string `json:"children" bson:"children"`
// 	XPos       float32  `json:"x_pos" bson:"x_pos"`
// 	YPos       float32  `json:"y_pos" bson:"y_pos"`
// }
