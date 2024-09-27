package register

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/authenticator/providers"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()

	handler := providers.Handler{
		Service: Model{},
	}

	router.Handle("/auth/google/register", http.HandlerFunc(handler.Auth))
	router.Handle("/auth/google/callback", http.HandlerFunc(handler.Callback))
	router.Handle("/auth/google/logout", http.HandlerFunc(handler.Logout))

	return router, nil
}
