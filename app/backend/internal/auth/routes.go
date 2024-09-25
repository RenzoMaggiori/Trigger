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
	router.Handle("POST /verify", http.HandlerFunc(handler.Verify))

	return router, nil
}
