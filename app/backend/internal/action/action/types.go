package action

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ActionCtxKey string = "ActionCtxKey"

type Service interface {
	Get() ([]ActionModel, error)
	GetById(primitive.ObjectID) (*ActionModel, error)
	GetByProvider(string) ([]ActionModel, error)
	GetByAction(string) (*ActionModel, error)
	Add(*AddActionModel) (*ActionModel, error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type ActionModel struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Input  []string           `json:"input" bson:"input"`
	Output []string           `json:"output" bson:"output"`
	// provider name (gmail, discord, github, ...)
	Provider string `json:"provider" bson:"provider"`
	// "trigger" or "action"
	Type string `json:"type" bson:"type"`
	// what does the action do? (send-email, delete-email, watch-push, ...)
	Action string `json:"action" bson:"action"`
}

type AddActionModel struct {
	Input    []string `json:"input" bson:"input"`
	Output   []string `json:"output" bson:"output"`
	Provider string   `json:"provider" bson:"provider"`
	Type     string   `json:"type" bson:"type"`
	Action   string   `json:"action" bson:"action"`
}
