package credentials

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()
	handler := Handler{
		Service: Model{},
	}

	server.Handle("POST /login", http.HandlerFunc(handler.Login))
	server.Handle("POST /register", http.HandlerFunc(handler.Register))
	server.Handle("POST /verify", http.HandlerFunc(handler.Verify))

	return router.NewRouter("/api/auth", server), nil
}
