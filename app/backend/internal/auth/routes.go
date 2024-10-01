package auth

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

	server.Handle("POST /login", http.HandlerFunc(handler.Login))
	server.Handle("POST /register", http.HandlerFunc(handler.Register))
	server.Handle("POST /verify", http.HandlerFunc(handler.Verify))

	return router.NewRouter("/auth", server), nil
}
