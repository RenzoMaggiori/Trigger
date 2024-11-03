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

	server.Handle("POST /create_pull_request", middlewares(http.HandlerFunc(handler.CreatePullRequest)))

	return router.NewRouter("/api/bitbucket/reaction", server), nil
}
