package twitch

import (
	"fmt"
	"os"

	gothicTwitch "github.com/markbates/goth/providers/twitch"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

func Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		Scopes: []string{
			gothicTwitch.ScopeUserReadEmail,
			gothicTwitch.ScopeModeratorReadFollowers,
			"user:write:chat",
		},
		Endpoint:    twitch.Endpoint,
		RedirectURL: fmt.Sprintf("http://localhost:%s/api/oauth2/callback", os.Getenv("AUTH_PORT")),
	}
}
