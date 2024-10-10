package providers

import (
	"time"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/pkg/authenticator"
)

type Service interface {
	authenticator.Authenticator
	Callback(user goth.User) (string, error)
}

type Handler struct {
	Service
}

type Model struct {
}

type ProviderCtx string

const AuthorizationHeaderCtxKey ProviderCtx = ProviderCtx("AuthorizationHeaderCtxKey")
const UserCtxKey ProviderCtx = ProviderCtx("UserCtxKey")
const LoginCtxKey ProviderCtx = ProviderCtx("LoginCtxKey")

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
