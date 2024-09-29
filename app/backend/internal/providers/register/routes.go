package register

import (
	"context"
	"net/http"

	"trigger.com/trigger/pkg/authenticator/providers"
	"trigger.com/trigger/pkg/router"
)

// This file function is to declare a global entrypoint for all the providers when registering a new User.
// Goth, the Auth library, will redirect each call to each providers callback so there no need to redeclare
// the same route in all the providers.

func Router(ctx context.Context) (router.PrefixedRouter, error) {
	server := http.NewServeMux()

	server.Handle("GET /register", http.HandlerFunc(providers.Auth))
	return router.NewPrefixedRouter("/provider", server), nil
}
