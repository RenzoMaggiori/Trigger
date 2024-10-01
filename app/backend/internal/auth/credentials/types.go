package credentials

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/authenticator"
)

type Service interface {
	authenticator.Authenticator
	Register(RegisterModel) (string, error)
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

var CredentialsCtxKey = "CredentialsCtxKey"

type CredentialsModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	User user.AddUserModel `json:"user"`
}
