package reaction

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

	server.Handle("POST /send_chat_message", middlewares(http.HandlerFunc(handler.SendChatMessage)))

	return router.NewRouter("/api/twitch/reaction", server), nil
}
