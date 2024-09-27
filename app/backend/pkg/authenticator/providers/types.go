package providers

import "context"

type Service interface {
	Auth(ctx context.Context) (string, error)
	Callback(ctx context.Context) (string, error)
	Logout(ctx context.Context) (string, error)
}

type Handler struct {
	Service
}

type Model struct {
}
