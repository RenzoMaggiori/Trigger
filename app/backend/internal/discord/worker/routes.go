package worker

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

	server.Handle("GET /me", middlewares(http.HandlerFunc(handler.Me)))
	server.Handle("GET /guilds/{guild_id}/channels", middlewares(http.HandlerFunc(handler.GetGuildChannels)))

	return router.NewRouter("/discord", server), nil
}
