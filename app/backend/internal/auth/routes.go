package auth

import (
	"context"
	"net/http"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()
	handler := Handler{
		Service: Model{},
	}

	router.Handle("POST /login", http.HandlerFunc(handler.Login))
	router.Handle("POST /register", http.HandlerFunc(handler.Register))

	return router, nil
}
