package router

import (
	"context"
	"net/http"
)

type RouterFn func(context.Context) (*http.ServeMux, error)

func Create(ctx context.Context, routers ...RouterFn) (*http.ServeMux, error) {
	router := http.NewServeMux()

	for _, routerFn := range routers {
		r, err := routerFn(ctx)
		if err != nil {
			return nil, err
		}

		router.Handle("/api/", http.StripPrefix("/api", r))
	}
	return router, nil
}
