package auth

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/mongodb"
)

var (
	errDatabaseNotFound error = errors.New("could not find mongo database")
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	database, ok := ctx.Value(mongodb.CtxKey).(*mongo.Database)
	if !ok {
		return nil, errDatabaseNotFound
	}

	router := http.NewServeMux()
	handler := Handler{
		Service: Model{
			DB:       database,
			authType: Undefined,
		},
	}

	router.Handle("POST /login", http.HandlerFunc(handler.Login))
	router.Handle("POST /register", http.HandlerFunc(handler.Register))
	router.Handle("POST /verify", http.HandlerFunc(handler.Verify))

	return router, nil
}
