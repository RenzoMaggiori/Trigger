package todo

import "net/http"

func Router() *http.ServeMux {
	router := http.NewServeMux()

	// router.Handle("/", )
	return router
}
