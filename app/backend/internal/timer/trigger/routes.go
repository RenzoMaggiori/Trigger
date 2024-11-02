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

	server.Handle("POST /watch_minute", middlewares(http.HandlerFunc(handler.WatchTimer)))
	server.Handle("POST /watch_hour", middlewares(http.HandlerFunc(handler.WatchTimer)))
	server.Handle("POST /watch_day", middlewares(http.HandlerFunc(handler.WatchTimer)))
	server.Handle("POST /webhook", middlewares(http.HandlerFunc(handler.WebhookTimer)))
	server.Handle("POST /stop", middlewares(http.HandlerFunc(handler.StopTimer)))

	return router.NewRouter("/api/timer/trigger", server), nil
}
