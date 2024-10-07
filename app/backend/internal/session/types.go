package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Get() ([]SessionModel, error)
	GetById(primitive.ObjectID) (*SessionModel, error)
	GetByUserId(primitive.ObjectID) ([]SessionModel, error)
	Add(*AddSessionModel) (*SessionModel, error)
	UpdateById(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error)
	DeleteById(primitive.ObjectID) error
	GetByToken(string) (*SessionModel, error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type SessionModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProviderName *string            `json:"provider_name,omitempty" bson:"provider_name,omitempty"`
	AccessToken  string             `json:"access_token" bson:"access_token"`
	RefreshToken *string            `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"id_token,omitempty" bson:"id_token,omitempty"`
}

type AddSessionModel struct {
	UserId       primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"access_token" bson:"access_token"`
	RefreshToken *string            `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"id_token,omitempty" bson:"id_token,omitempty"`
}

type UpdateSessionModel struct {
	AccessToken  *string    `json:"access_token,omitempty" bson:"access_token,omitempty"`
	RefreshToken *string    `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
	IdToken      *string    `json:"id_token,omitempty" bson:"id_token,omitempty"`
}
