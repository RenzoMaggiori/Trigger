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

	server.Handle("POST /watch", middlewares(http.HandlerFunc(handler.Watch)))
	server.Handle("POST /webhook", http.HandlerFunc(handler.Webhook))
	server.Handle("POST /stop", middlewares(http.HandlerFunc(handler.Stop)))

	return router.NewRouter("/twitch/trigger", server), nil
}
