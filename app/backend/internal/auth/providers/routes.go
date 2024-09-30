package providers

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (router.Router, error) {
	server := http.NewServeMux()

	handler := Handler{
		Service: Model{},
	}

	server.Handle("GET /login", http.HandlerFunc(handler.Login))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /logout", http.HandlerFunc(handler.Logout))
	return router.NewRouter("/oauth2", server), nil
}
