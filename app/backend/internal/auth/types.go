package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/auth/authenticator"
)

type Service interface {
	authenticator.Authenticator
	GetToken(string) (string, error)
	VerifyToken(string) error
}

type Handler struct {
	Service
}

type AuthType int64

const (
	Undefined AuthType = iota
	Credentials
	OAuth
)

type Model struct {
	DB       *mongo.Database
	authType AuthType
}

type CredentialsCtx string

const CredentialsCtxKey CredentialsCtx = CredentialsCtx("CredentialsCtxKey")

type CredentialsModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	User user.AddUserModel `json:"user"`
}
