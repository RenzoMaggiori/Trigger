package middleware

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		if origin == "http://localhost:3000" {
			res.Header().Add("Access-Control-Allow-Origin", origin)
		}

		res.Header().Add("Access-Control-Allow-Credentials", "true")
		res.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		res.Header().Add("Access-Control-Expose-Headers", "Authorization")
		res.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if req.Method == "OPTIONS" {
			http.Error(res, "No Content", http.StatusNoContent)
			return
		}
		next.ServeHTTP(res, req)
	})
}
