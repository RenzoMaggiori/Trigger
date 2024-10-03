package action

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ActionCtxKey string = "ActionCtxKey"

type Service interface {
	Get() ([]ActionModel, error)
	GetById(primitive.ObjectID) (*ActionModel, error)
	GetByActionType(actionType string) ([]ActionModel, error)
	Add(*AddActionModel) (*ActionModel, error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type ActionType int

const (
	Trigger ActionType = iota
	Action
)

var ActionTypeValue = map[string]ActionType{
	"trigger": Trigger,
	"action":  Action,
}

type ActionModel struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Input      []string           `json:"input" bson:"input"`
	Output     []string           `json:"output" bson:"output"`
	ActionType string             `json:"action_type" bson:"action_type"`
}

type AddActionModel struct {
	Input      []string `json:"input" bson:"input"`
	Output     []string `json:"output" bson:"output"`
	ActionType string   `json:"action_type" bson:"action_type"`
}
