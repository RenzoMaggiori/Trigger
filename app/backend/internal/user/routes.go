package user

import (
	"context"
	"net/http"
)

func Router(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()
	return router, nil
}
