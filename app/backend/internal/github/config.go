package github

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_KEY"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
		RedirectURL:  fmt.Sprintf("http://localhost:%s/api/oauth2/callback", os.Getenv("AUTH_PORT")),
	}
}
