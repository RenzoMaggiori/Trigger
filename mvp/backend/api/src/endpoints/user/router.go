package user

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/api/src/database"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	database, ok := ctx.Value(database.CtxKey).(*mongo.Client)
	if !ok {
		return nil, fmt.Errorf("could not get Database from Context")
	}

	router := http.NewServeMux()
	handler := Handler{
		Service: Model{Mongo: database},
	}

	router.Handle("GET /user/{email}", http.HandlerFunc(handler.GetByEmail))
	router.Handle("POST /user", http.HandlerFunc(handler.Add))
	return router, nil
}
