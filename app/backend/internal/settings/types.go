package settings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/sync"
)

type Service interface {
	GetById(primitive.ObjectID) (*sync.SyncModel, error)
	GetByUserId(string) (*sync.SyncModel, error)
	// Add(*AddSettingsModel) (*SettingsModel, error)
	// UpdateById(primitive.ObjectID, *UpdateSettingsModel) (*SettingsModel, error)
	// UpdateByEmail(string, *UpdateSettingsModel) (*SettingsModel, error)
	// DeleteById(primitive.ObjectID) error
	// DeleteByEmail(string) error
}

type Handler struct {
	Service
}

type Model struct {
}

type SettingsModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type AddSettingsModel struct {
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	ProviderName *string            `json:"providerName,omitempty" bson:"providerName,omitempty"`
	AccessToken  string             `json:"accessToken" bson:"accessToken"`
	RefreshToken *string            `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       time.Time          `json:"expiry" bson:"expiry"`
	IdToken      *string            `json:"idToken,omitempty" bson:"idToken,omitempty"`
}

type UpdateSettingsModel struct {
	AccessToken  *string    `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken *string    `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
	IdToken      *string    `json:"idToken,omitempty" bson:"idToken,omitempty"`
}