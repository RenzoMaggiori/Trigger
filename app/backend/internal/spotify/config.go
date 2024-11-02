package spotify

import (
	"fmt"
	"os"

	gothicSpotify "github.com/markbates/goth/providers/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

func Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_KEY"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		Scopes: []string{
			gothicSpotify.ScopeUserReadEmail,
			gothicSpotify.ScopeUserReadPrivate,
			gothicSpotify.ScopeUserReadPlaybackState,
			gothicSpotify.ScopeUserModifyPlaybackState,
		},
		Endpoint:    spotify.Endpoint,
		RedirectURL: fmt.Sprintf("http://localhost:%s/api/oauth2/callback", os.Getenv("AUTH_PORT")),
	}
}
