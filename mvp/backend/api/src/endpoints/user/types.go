package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email         string    `json:"email" bson:"email"`
	AccessToken   string    `json:"accessToken" bson:"access_token"`
	RefreshToken  string    `json:"refreshToken" bson:"refresh_token,omitempty"`
	TokenType     string    `json:"tokenType" bson:"token_type,omitempty"`
	Expiry        time.Time `json:"expiry" bson:"expiry,omitempty"`
	LastHistoryId int64     `json:"lastHistoryId" bson:"last_history_id,omitempty"`
}

type UpdateUser struct {
	LastHistoryId int64 `json:"lastHistoryId" bson:"last_history_id,omitempty"`
}

type Service interface {
	Add(User) (*primitive.ObjectID, error)
	UpdateByEmail(string, UpdateUser) (*primitive.ObjectID, error)
	GetByEmail(string) (*User, error)
}

type Handler struct {
	Service
}

type Model struct {
	Mongo *mongo.Client
}
