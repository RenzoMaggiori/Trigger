package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
)

var AuthHeaderCtxKey string = "authHeaderCtxKey"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accessHeader := strings.Split(req.Header.Get("Authorization"), " ")
		if len(accessHeader) < 2 {
			log.Println("could not retrieve access token from header")
			http.Error(res, "could not retrieve access token from header", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), AuthHeaderCtxKey, accessHeader[1])
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
