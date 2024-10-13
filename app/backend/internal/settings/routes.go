package settings

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
)

var (
	// errDatabaseNotFound error = errors.New("could not find mongo database")
)

func Router(ctx context.Context) (*router.Router, error) {
	settingsCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errors.New("could not find sync mongo collection")
	}

	server := http.NewServeMux()

	handler := Handler{
		Service: Model{
			Collection: settingsCollection,
		},
	}
	server.Handle("GET /id/{id}", http.HandlerFunc(handler.GetById))
	server.Handle("POST /add/{id}", http.HandlerFunc(handler.Add))
	server.Handle("GET /user_id/{id}", http.HandlerFunc(handler.GetByUserId))

	return router.NewRouter("/settings", server), nil
}
