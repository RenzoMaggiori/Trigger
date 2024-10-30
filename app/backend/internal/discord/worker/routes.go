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
	server.Handle("GET /current", middlewares(http.HandlerFunc(handler.GetCurrentDiscordSession)))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.AddDiscordSession)))
	server.Handle("PUT /update/{user_id}", middlewares(http.HandlerFunc(handler.UpdateDiscordSession)))
	server.Handle("POST /delete/{user_id}", middlewares(http.HandlerFunc(handler.DeleteDiscordSession)))

	return router.NewRouter("/discord/worker", server), nil
}
