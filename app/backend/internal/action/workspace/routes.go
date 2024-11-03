package workspace

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/router"
)

var (
	errCollectionNotFound error = errors.New("could not find session mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	UserActionCollection, ok := ctx.Value(WorkspaceCtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: UserActionCollection,
		},
	}

	server.Handle("GET /", middlewares(http.HandlerFunc(handler.Get)))
	server.Handle("GET /me", middlewares(http.HandlerFunc(handler.GetByAcessToken)))
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetById)))
	server.Handle("GET /user_id/{user_id}", middlewares(http.HandlerFunc(handler.GetByUserId)))
	server.Handle("GET /action_id/{action_id}", middlewares(http.HandlerFunc(handler.GetByActionId)))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.Add)))
	server.Handle("PATCH /id/{id}", middlewares(http.HandlerFunc(handler.UpdateById)))
	server.Handle("PATCH /start/id/{id}", middlewares(http.HandlerFunc(handler.Start)))
	server.Handle("PATCH /stop/id/{id}", middlewares(http.HandlerFunc(handler.Stop)))
	server.Handle("PATCH /action_completed", middlewares(http.HandlerFunc(handler.ActionCompleted)))
	server.Handle("PATCH /watch_completed", middlewares(http.HandlerFunc(handler.WatchCompleted)))
	server.Handle("DELETE /id/{id}", middlewares(http.HandlerFunc(handler.DeleteById)))
	server.Handle("GET /templates", middlewares(http.HandlerFunc(handler.GetTemplates)))

	return router.NewRouter("/api/workspace", server), nil
}
