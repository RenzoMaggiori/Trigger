package router

import (
	"net/http"

	"trigger.com/api/src/endpoints/todo"
)

func Create() *http.ServeMux {
	router := http.NewServeMux()
	routers := []*http.ServeMux{
		todo.Router(),
	}

	for _, r := range routers {
		router.Handle("/api/", http.StripPrefix("/api", r))
	}
	return router
}
