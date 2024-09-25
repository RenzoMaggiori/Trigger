package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	oauthprovider "trigger.com/trigger/pkg/auth/oauth-provider"
)

func config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
		RedirectURL:  fmt.Sprintf("%s/api/callback", os.Getenv("GITHUB_SERVICE_BASE_URL")),
	}
}

func Router(ctx context.Context) (*http.ServeMux, error) {
	router := http.NewServeMux()

	handler := Handler{
		Service: Model{
			OAuth2Provider: oauthprovider.New(config()),
		},
	}

	router.Handle("POST /login", http.HandlerFunc(handler.GithubLogin))
	router.Handle("POST /callback", http.HandlerFunc(handler.GithubCallback))

	return router, nil
}
