package credentials

import (
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/auth/authenticator"
)

type Service interface {
	authenticator.Authenticator
	Register(RegisterModel) (string, error)
	VerifyToken(string) error
}

type Handler struct {
	Service
}

type Model struct {
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
