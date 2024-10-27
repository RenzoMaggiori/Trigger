package settings

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GetById(primitive.ObjectID) (*SettingsResponseModel, error)
	GetByUserId(primitive.ObjectID) ([]SettingsResponseModel, error)
	Add(*AddSettingsModel) (error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type SettingsModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	Active       bool               `json:"active" bson:"active"`
}

type AddSettingsModel struct {
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	Active       bool               `json:"active" bson:"active"`
}

type UpdateSettingsModel struct {
	Active       bool               `json:"active" bson:"active"`
}

type SettingsResponseModel struct {
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	Active       bool               `json:"active" bson:"active"`
}