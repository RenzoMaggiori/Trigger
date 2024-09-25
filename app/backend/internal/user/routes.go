package user

import (
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errCollectionNotFound error = errors.New("could not find user mongo colleciton")
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	userCollection, ok := ctx.Value(UserCollectionCtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	}

	router := http.NewServeMux()
	handler := Handler{
		Service: Model{
			Collection: userCollection,
		},
	}

	// TODO: handle auth
	router.Handle("GET /user", http.HandlerFunc(handler.GetUsers))
	router.Handle("GET /user/{id}", http.HandlerFunc(handler.GetUserById))
	router.Handle("GET /user/{email}", http.HandlerFunc(handler.GetUserByEmail))
	router.Handle("POST /user", http.HandlerFunc(handler.AddUser))
	router.Handle("PATCH /user/{id}", http.HandlerFunc(handler.UpdateUserById))
	router.Handle("PATCH /user/{email}", http.HandlerFunc(handler.UpdateUserByEmail))
	router.Handle("DELETE /user/{id}", http.HandlerFunc(handler.DeleteUserById))
	router.Handle("DELETE /user/{email}", http.HandlerFunc(handler.DeleteUserById))

	return router, nil
}
