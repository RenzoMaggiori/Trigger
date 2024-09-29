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
		router.Handle(r.Prefix,
			http.StripPrefix(r.Prefix, r.Server))
	}

	return router, nil
}
