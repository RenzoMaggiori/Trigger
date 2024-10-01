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
	errCollectionNotFound error = errors.New("could not find user mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	userCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	server := http.NewServeMux()
	// middlewares := middleware.Create(
	// 	middleware.Auth,
	// )
	handler := Handler{
		Service: Model{
			Collection: userCollection,
		},
	}

	server.Handle("GET /", http.HandlerFunc(handler.GetSessions))
	server.Handle("GET /id/{id}", http.HandlerFunc(handler.GetSessionById))
	server.Handle("GET /userId/{userId}", http.HandlerFunc(handler.GetSessionByUserId))
	server.Handle("POST /add", http.HandlerFunc(handler.AddSession))
	server.Handle("PATCH /id/{id}", http.HandlerFunc(handler.UpdateSessionById))
	server.Handle("PATCH /userId/{userId}", http.HandlerFunc(handler.UpdateSessionByUserId))
	server.Handle("DELETE /id/{id}", http.HandlerFunc(handler.DeleteSessionById))
	server.Handle("DELETE /userId/{userId}", http.HandlerFunc(handler.DeleteSessionByUserId))

	return router.NewRouter("/user", server), nil
}
