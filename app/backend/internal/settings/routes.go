package settings

import (
	"context"
	"errors"
	"net/http"

	"trigger.com/trigger/pkg/router"
)

var (
	errDatabaseNotFound error = errors.New("could not find mongo database")
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()
	handler := Handler{
		Service: Model{},
	}
	server.Handle("GET /id/{id}", http.HandlerFunc(handler.ById))
	server.Handle("GET /user_id/{id}", http.HandlerFunc(handler.GetByUserId))

	return router.NewRouter("/settings", server), nil
}
