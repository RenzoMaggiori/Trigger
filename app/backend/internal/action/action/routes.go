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

	server.Handle("GET /about.json", http.HandlerFunc(handler.About))
	server.Handle("GET /api/action/", http.HandlerFunc(handler.GetActions))
	server.Handle("GET /api/action/id/{id}", http.HandlerFunc(handler.GetActionById))
	server.Handle("GET /api/action/provider/{provider}", http.HandlerFunc(handler.GetActionsByProvider))
	server.Handle("GET /api/action/action/{action}", http.HandlerFunc(handler.GetActionByAction))
	server.Handle("POST /api/action/add", http.HandlerFunc(handler.AddAction))
	return router.NewRouter("", server), nil
}
