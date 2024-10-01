package providers

import (
	"time"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/authenticator"
)

type Service interface {
	authenticator.Authenticator
	Callback(user goth.User) (string, error)
	// Get() ([]SessionModel, error)
	// GetById(primitive.ObjectID) (*SessionModel, error)
	// GetByUserId(primitive.ObjectID) (*SessionModel, error)
	// Add(*AddSessionModel) (*SessionModel, error)
	// UpdateById(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error)
	// UpdateByUserId(primitive.ObjectID, *UpdateSessionModel) (*SessionModel, error)
	// DeleteById(primitive.ObjectID) error
	// DeleteByUserId(primitive.ObjectID) error
}

type Handler struct {
	Service
}

type Model struct {
	DB *mongo.Database
}

var CredentialsCtxKey = "CredentialsCtxKey"

type SessionModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type AddSessionModel struct {
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type UpdateSessionModel struct {
	AccessToken  *string    `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken *string    `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
	IdToken      *string    `json:"idToken,omitempty" bson:"idToken,omitempty"`
}
