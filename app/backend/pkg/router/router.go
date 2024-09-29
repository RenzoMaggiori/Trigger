package router

import (
	"context"
	"net/http"
)

type RouterFn func(context.Context) (PrefixedRouter, error)

type PrefixedRouter struct {
	Server *http.ServeMux
	Prefix string
}

func NewPrefixedRouter(prefix string, server *http.ServeMux) PrefixedRouter {
	return PrefixedRouter{
		Server: server,
		Prefix: prefix,
	}
}

func Create(ctx context.Context, routers ...RouterFn) (*http.ServeMux, error) {
	router := http.NewServeMux()

	for _, prefixedRouter := range routers {
		r, err := prefixedRouter(ctx)
		if err != nil {
			return nil, err
		}
		prefix := "/api" + r.Prefix
		router.Handle(prefix+"/",
			http.StripPrefix(prefix, r.Server))
	}

	return router, nil
}
