package todo

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"trigger.com/api/src/database"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	database, ok := ctx.Value(database.CtxKey).(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("Could not get Database from Context")
	}

	router := http.NewServeMux()
	handler := Handler{Todos: Model{
		db: database,
	}}

	router.HandleFunc("GET /todo", handler.GetAll)
	router.HandleFunc("GET /todo/{id}", handler.GetById)
	router.HandleFunc("POST /todo", handler.Add)
	router.HandleFunc("PATCH /todo/{id}", handler.Patch)
	router.HandleFunc("DELETE /todo/{id}", handler.Delete)
	return router, nil
}
