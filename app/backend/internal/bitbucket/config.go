package bitbucket

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/bitbucket"
)

func Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("BITBUCKET_KEY"),
		ClientSecret: os.Getenv("BITBUCKET_SECRET"),
		Scopes:       []string{},
		Endpoint:     bitbucket.Endpoint,
		RedirectURL:  fmt.Sprintf("%s/api/oauth2/callback", os.Getenv("SERVER_BASE_URL")),
	}
}
