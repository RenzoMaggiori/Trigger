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

	server.Handle("GET /", middlewares(http.HandlerFunc(handler.GetWorkspace)))
	server.Handle("GET /me", middlewares(http.HandlerFunc(handler.GetMyWorkspaces)))
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetWorkspaceById)))
	server.Handle("GET /user_id/{user_id}", middlewares(http.HandlerFunc(handler.GetWorkspacesByUserId)))
	server.Handle("GET /action_id/{action_id}", middlewares(http.HandlerFunc(handler.GetWorkspacesByActionId)))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.AddWorkspace)))
	server.Handle("PATCH /action_completed", middlewares(http.HandlerFunc(handler.ActionCompletedWorkspace)))
	server.Handle("PATCH /id/{id}", middlewares(http.HandlerFunc(handler.UpdateById)))
	server.Handle("PATCH /watch_completed", middlewares(http.HandlerFunc(handler.WatchCompleted)))
	// server.Handle("DELETE /id/{id}", middlewares(http.HandlerFunc(handler.DeleteById)))

	return router.NewRouter("/workspace", server), nil
}
