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
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetWorkspaceById)))
	server.Handle("GET /user_id/{user_id}", middlewares(http.HandlerFunc(handler.GetWorkspacesByUserId)))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.AddWorkspace)))
	server.Handle("PATCH /completed_action/{user_id}", middlewares(http.HandlerFunc(handler.UpdateActionCompletedWorkspace)))
	// server.Handle("PATCH /id/{id}", middlewares(http.HandlerFunc(handler.UpdateById)))
	// server.Handle("DELETE /id/{id}", middlewares(http.HandlerFunc(handler.DeleteById)))

	return router.NewRouter("/workspace", server), nil
}
