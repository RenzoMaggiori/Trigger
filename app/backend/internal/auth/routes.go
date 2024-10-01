package auth

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
)

var (
	errDatabaseNotFound error = errors.New("could not find mongo database")
)

func Router(ctx context.Context) (*router.Router, error) {
	database, ok := ctx.Value(mongodb.CtxKey).(*mongo.Database)
	if !ok {
		return nil, errDatabaseNotFound
	}

	server := http.NewServeMux()
	handler := Handler{
		Service: Model{
			DB:       database,
			authType: Undefined,
		},
	}

	server.Handle("POST /login", http.HandlerFunc(handler.Login))
	server.Handle("POST /register", http.HandlerFunc(handler.Register))
	server.Handle("POST /verify", http.HandlerFunc(handler.Verify))

	return router.NewRouter("/auth", server), nil
}
