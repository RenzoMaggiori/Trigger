package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/jwt"
)

type Handler struct {
	secret string
}

type TokenCtx string

const TokenCtxKey = TokenCtx("TokenCtxKey")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := fetch.Fetch(
			&http.Client{},
			fetch.NewFetchRequest(
				http.MethodPost,
				fmt.Sprintf("%s/api/auth/verify", os.Getenv("AUTH_SERVICE_BASE_URL")),
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

		token, err := jwt.FromRequest(r.Header.Get("Authorization"))
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		log.Println("middleware: ", token)
		next.ServeHTTP(
			w,
			r.WithContext(
				context.WithValue(
					r.Context(), TokenCtxKey, token,
				),
			),
		)
	})
}
