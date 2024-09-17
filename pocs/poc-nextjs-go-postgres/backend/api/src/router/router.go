package router

import (
	"context"
	"net/http"

	"trigger.com/api/src/endpoints/todo"
)

type RouterFn func(context.Context) (*http.ServeMux, error)

func Create(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()
	routers := []RouterFn{
		todo.Router,
	}

	for _, routerFn := range routers {
		r, err := routerFn(ctx)
		if err != nil {
			return nil, err
		}

		router.Handle("/api/", http.StripPrefix("/api", r))
	}
	return router, nil
}
