package worker

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

	server.Handle("GET /me", middlewares(http.HandlerFunc(handler.Me)))
	server.Handle("GET /guilds/{guild_id}/channels", middlewares(http.HandlerFunc(handler.GetGuildChannels)))

	return router.NewRouter("/discord", server), nil
}
