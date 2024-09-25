package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/fetch"
)

type Handler struct {
	secret string
}

type TokenCtx string

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := fetch.Fetch(
			&http.Client{},
			fetch.NewFetchRequest(
				http.MethodPost,
				fmt.Sprintf("%s/api/verify", os.Getenv("AUTH_SERVICE_BASE_URL")),
				nil,
				map[string]string{
					"Authorization": r.Header.Get("Authorization"),
				},
			),
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if res.StatusCode != http.StatusOK {
			log.Printf("invalid status code, received %s\n", res.Status)
			http.Error(w, "could not verify authorization", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
