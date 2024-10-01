package session

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
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
	// middlewares := middleware.Create(
	// 	middleware.Auth,
	// )
	handler := Handler{
		Service: Model{
			Collection: sessionCollection,
		},
	}

	server.Handle("GET /", http.HandlerFunc(handler.GetSessions))
	server.Handle("GET /id/{id}", http.HandlerFunc(handler.GetSessionById))
	server.Handle("GET /user_id/{user_id}", http.HandlerFunc(handler.GetSessionByUserId))
	server.Handle("GET /access_token/{access_token}", http.HandlerFunc(handler.GetByToken))
	server.Handle("POST /add", http.HandlerFunc(handler.AddSession))
	server.Handle("PATCH /id/{id}", http.HandlerFunc(handler.UpdateSessionById))
	server.Handle("DELETE /id/{id}", http.HandlerFunc(handler.DeleteSessionById))

	return router.NewRouter("/session", server), nil
}
