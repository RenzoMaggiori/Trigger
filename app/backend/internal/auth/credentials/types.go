package credentials

import (
	"context"

	"trigger.com/trigger/internal/user"
)

type Service interface {
	Login(CredentialsModel) (string, error)
	Register(RegisterModel) (string, error)
	VerifyToken(string) error
	Logout(ctx context.Context) error
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
