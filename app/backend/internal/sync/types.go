package sync

import (
	"net/http"
	"time"

	"github.com/markbates/goth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GrantAccess(w http.ResponseWriter, r *http.Request) error
	SyncWith(gothUser goth.User, access_token string) (error)
	Callback(gothUser goth.User, access_token string) (error)
	ByUserId(userId primitive.ObjectID, provider string) (SyncModel, error)
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type SyncModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type AddSyncModel struct {
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type UpdateSyncModel struct {
	AccessToken  *string    `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken *string    `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
	IdToken      *string    `json:"idToken,omitempty" bson:"idToken,omitempty"`
}
