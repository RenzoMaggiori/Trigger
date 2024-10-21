package session

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
)

var (
	errCollectionNotFound error = errors.New("could not find session mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	sessionCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: sessionCollection,
		},
	}

	server.Handle("GET /", middlewares(http.HandlerFunc(handler.GetSessions)))
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetSessionById)))
	server.Handle("GET /user_id/{user_id}", middlewares(http.HandlerFunc(handler.GetSessionByUserId)))
	server.Handle("GET /access_token/{access_token}", http.HandlerFunc(handler.GetByToken))
	server.Handle("POST /add", middlewares(http.HandlerFunc(handler.AddSession)))
	server.Handle("PATCH /id/{id}", middlewares(http.HandlerFunc(handler.UpdateSessionById)))
	server.Handle("DELETE /id/{id}", middlewares(http.HandlerFunc(handler.DeleteSessionById)))

	return router.NewRouter("/session", server), nil
}
