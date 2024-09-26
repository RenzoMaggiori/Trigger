package google

import (
	"context"
	"net/http"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()
	handler := Handler{
		Service: Model{},
	}

	router.Handle("GET /auth/google/auth", http.HandlerFunc(handler.Auth))
	router.Handle("GET /auth/gmail/callback", http.HandlerFunc(handler.Callback))
	router.Handle("GET /auth/google/logout", http.HandlerFunc(handler.Logout))

	return router, nil
}
