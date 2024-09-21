package gmail

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
		return nil, fmt.Errorf("Could not get Database from Context")
	}
	_ = database

	router := http.NewServeMux()
	// handler := Handler{Gmail: }

	return router, nil
}
