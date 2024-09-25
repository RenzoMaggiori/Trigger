package auth

import (
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
}

var CredentialsCtxKey = "CredentialsCtxKey"

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterModel struct {
	User user.AddUserModel `json:"user"`
}
