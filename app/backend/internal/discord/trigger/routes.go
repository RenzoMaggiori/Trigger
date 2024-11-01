package trigger

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	discordCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errors.New("could not find discord mongo collection")
	}
	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: &Model{
			Collection: discordCollection,
		},
	}

	server.Handle("POST /watch_channel_message", middlewares((http.HandlerFunc(handler.WatchDiscord))))
	server.Handle("POST /stop", middlewares(http.HandlerFunc(handler.StopDiscord)))
	server.Handle("POST /webhook", middlewares(http.HandlerFunc(handler.WebhookDiscord)))

	server.Handle("GET /sessions", (http.HandlerFunc(handler.GetDiscordSessions)))

	return router.NewRouter("/api/discord/trigger", server), nil
}
