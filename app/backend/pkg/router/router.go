package router

import (
	"context"
	"net/http"
)

type RouterFn func(context.Context) (Router, error)

type Router struct {
	Server *http.ServeMux
	Prefix string
}

func NewRouter(prefix string, server *http.ServeMux) Router {
	return Router{
		Server: server,
		Prefix: prefix,
	}
}

func Create(ctx context.Context, routers ...RouterFn) (*http.ServeMux, error) {
	router := http.NewServeMux()

	for _, Router := range routers {
		r, err := Router(ctx)
		if err != nil {
			return nil, err
		}
		prefix := "/api" + r.Prefix
		router.Handle(prefix+"/",
			http.StripPrefix(prefix, r.Server))
	}

	return router, nil
}
