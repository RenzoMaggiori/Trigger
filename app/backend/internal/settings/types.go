package settings

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GetById(primitive.ObjectID) (*SettingsModel, error)
	GetByUserId(primitive.ObjectID) ([]SettingsModel, error)
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
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	Active       bool               `json:"active" bson:"active"`
}

type AddSettingsModel struct {
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	Active       bool               `json:"active" bson:"active"`
}

type UpdateSettingsModel struct {
	Active       bool               `json:"active" bson:"active"`
}
