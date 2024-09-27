package providers

import (
	"context"

	"github.com/markbates/goth"
)

type Service interface {
	Callback(user goth.User) (string, error)
	Logout(ctx context.Context) (string, error)
}

type Handler struct {
	Service
}
