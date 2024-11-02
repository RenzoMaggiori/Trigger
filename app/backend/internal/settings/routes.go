package settings

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
	settingsCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errors.New("could not find settings mongo collection")
	}

	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: settingsCollection,
		},
	}
	server.Handle("GET /me", middlewares(http.HandlerFunc(handler.GetMySettings)))
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetById)))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.Add)))
	server.Handle("GET /user_id/{id}", middlewares(http.HandlerFunc(handler.GetByUserId)))
	server.Handle("PATCH /update", middlewares(http.HandlerFunc(handler.Update)))

	return router.NewRouter("/api/settings", server), nil
}
