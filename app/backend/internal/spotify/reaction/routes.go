package reaction

import (
	"context"
	"errors"
	"net/http"

	/* "go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb" */
	"trigger.com/trigger/pkg/router"
)

var (
	errCollectionNotFound error = errors.New("could not find spotify mongo colleciton")
)

func Router(ctx context.Context) (*router.Router, error) {
	/* spotifyCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errCollectionNotFound
	} */

	server := http.NewServeMux()
	/* middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: userCollection,
		},
	}

	server.Handle("GET /", middlewares(http.HandlerFunc(handler.GetUsers)))
	server.Handle("GET /id/{id}", middlewares(http.HandlerFunc(handler.GetUserById)))
	server.Handle("GET /email/{email}", middlewares(http.HandlerFunc(handler.GetUserByEmail)))
	server.Handle("POST /add", http.HandlerFunc(handler.AddUser))
	server.Handle("PATCH /id/{id}", middlewares(http.HandlerFunc(handler.UpdateUserById)))
	server.Handle("PATCH /email/{email}", middlewares(http.HandlerFunc(handler.UpdateUserByEmail)))
	server.Handle("DELETE /id/{id}", middlewares(http.HandlerFunc(handler.DeleteUserById)))
	server.Handle("DELETE /email/{email}", middlewares(http.HandlerFunc(handler.DeleteUserById))) */

	return router.NewRouter("/spotify/reaction", server), nil

}
