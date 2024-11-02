package trigger

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {

	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{},
	}

	server.Handle("POST /watch_push", middlewares(http.HandlerFunc(handler.WatchGithub)))
	server.Handle("POST /stop", middlewares(http.HandlerFunc(handler.StopGithub)))
	server.Handle("POST /webhook", http.HandlerFunc(handler.WebhookGithub))

	return router.NewRouter("/api/github/trigger", server), nil
}
