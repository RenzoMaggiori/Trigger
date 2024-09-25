package user

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
)

var (
	errCollectionNotFound error = errors.New("could not find user mongo colleciton")
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	userCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	router := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: userCollection,
		},
	}

	router.Handle("GET /user", middlewares(http.HandlerFunc(handler.GetUsers)))
	router.Handle("GET /user/id/{id}", middlewares(http.HandlerFunc(handler.GetUserById)))
	router.Handle("GET /user/email/{email}", middlewares(http.HandlerFunc(handler.GetUserByEmail)))
	router.Handle("POST /user", http.HandlerFunc(handler.AddUser))
	router.Handle("PATCH /user/id/{id}", middlewares(http.HandlerFunc(handler.UpdateUserById)))
	router.Handle("PATCH /user/email/{email}", middlewares(http.HandlerFunc(handler.UpdateUserByEmail)))
	router.Handle("DELETE /user/id/{id}", middlewares(http.HandlerFunc(handler.DeleteUserById)))
	router.Handle("DELETE /user/email/{email}", middlewares(http.HandlerFunc(handler.DeleteUserById)))

	return router, nil
}
