package action

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/router"
)

var (
	errCollectionNotFound error = errors.New("could not find session mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	actionCollection, ok := ctx.Value(ActionCtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	server := http.NewServeMux()
	// middlewares := middleware.Create(
	// 	middleware.Auth,
	// )
	handler := Handler{
		Service: Model{
			Collection: actionCollection,
		},
	}

	server.Handle("GET /", http.HandlerFunc(handler.GetActions))
	server.Handle("GET /id/{id}", http.HandlerFunc(handler.GetActionById))
	server.Handle("GET /provider/{provider}", http.HandlerFunc(handler.GetActionsByProvider))
	server.Handle("GET /action/{action}", http.HandlerFunc(handler.GetActionByAction))
	server.Handle("POST /add", http.HandlerFunc(handler.AddAction))
	return router.NewRouter("/action", server), nil
}
