package trigger

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type StopModel struct {
	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	HookId string `json:"hookId"`
}
