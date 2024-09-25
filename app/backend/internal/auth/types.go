package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/authenticator"
)

type Service interface {
	authenticator.Authenticator
}

type Handler struct {
	Service
}

type Model struct {
	DB *mongo.Database
}

var CredentialsCtxKey = "CredentialsCtxKey"

type CredentialsModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	User user.AddUserModel `json:"user"`
}
