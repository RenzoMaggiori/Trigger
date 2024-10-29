package trigger

import (
	"context"
	"errors"
	"net/http"

	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
)

var (
	errCollectionNotFound error = errors.New("could not find spotify mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{},
	}

	server.Handle("POST /watch_followers", middlewares(http.HandlerFunc(handler.WatchSpotify)))
	server.Handle("POST /webhook", middlewares(http.HandlerFunc(handler.WebhookSpotify)))
	server.Handle("POST /stop", middlewares(http.HandlerFunc(handler.StopSpotify)))

	return router.NewRouter("/spotify/trigger", server), nil
}
