package google

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

	// Prefix routes with /google
	router.Handle("GET /google/register/callback", http.HandlerFunc(handler.Callback))
	router.Handle("GET /google/register/logout", http.HandlerFunc(handler.Logout))

	return router, nil
}
