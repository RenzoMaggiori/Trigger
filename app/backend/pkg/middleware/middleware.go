package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Create(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range middlewares {
			next = m(next)
		}
		return next
	}
}
