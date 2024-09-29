package github

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/authenticator/providers"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (router.PrefixedRouter, error) {
	server := http.NewServeMux()

	handler := providers.Handler{
		Service: Model{},
	}

	server.Handle("GET /register/callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /register/logout", http.HandlerFunc(handler.Logout))

	return router.NewPrefixedRouter("/github", server), nil
}
