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

	server.Handle("POST /watch", middlewares(http.HandlerFunc(handler.WatchGmail)))
	server.Handle("POST /webhook", http.HandlerFunc(handler.WebhookGmail))
	server.Handle("POST /stop", http.HandlerFunc(handler.StopGmail))

	return router.NewRouter("/gmail/trigger", server), nil
}
