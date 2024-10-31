package action

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ActionCtxKey string = "ActionCtxKey"

type Service interface {
	About(string) (AboutModel, error)
	Get() ([]ActionModel, error)
	GetById(primitive.ObjectID) (*ActionModel, error)
	GetByProvider(string) ([]ActionModel, error)
	GetByActionName(string) (*ActionModel, error)
	Add(*AddActionModel) (*ActionModel, error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type AboutModel struct {
	Client ClientModel `json:"client"`
	Server ServerModel `json:"server"`
}

type ClientModel struct {
	Host string `json:"host"`
}

type ServerModel struct {
	CurrentTime int64          `json:"current_time"`
	Services    []ServiceModel `json:"services"`
}

type ServiceModel struct {
	Name      string      `json:"name"`
	Actions   []AreaModel `json:"actions"`
	Reactions []AreaModel `json:"reactions"`
}

type AreaModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ActionModel struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Input  []string           `json:"input" bson:"input"`
	Output []string           `json:"output" bson:"output"`
	// provider name (gmail, discord, github, ...)
	Provider string `json:"provider" bson:"provider"`
	// "trigger" or "reaction"
	Type string `json:"type" bson:"type"`
	// what does the action do? (send-email, delete-email, watch-push, ...)
	Action string `json:"action" bson:"action"`
}

type AddActionModel struct {
	Provider string   `json:"provider" bson:"provider"`
	Type     string   `json:"type" bson:"type"`
	Action   string   `json:"action" bson:"action"`
	Input    []string `json:"input" bson:"input"`
	Output   []string `json:"output" bson:"output"`
}
