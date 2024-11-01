package providers

import (
	"context"
	"net/http"
	"os"

	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/twitch"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()

	handler := Handler{
		Service: Model{},
	}

	callback := os.Getenv("AUTH_CALLBACK_URL")
	CreateProvider(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			callback,
			"https://mail.google.com/",
			"https://www.googleapis.com/auth/gmail.send",
			"email",
		),
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callback,
			"repo",
			"write:repo_hook",
		),
		discord.New(
			os.Getenv("DISCORD_KEY"),
			os.Getenv("DISCORD_SECRET"),
			callback,
			discord.ScopeIdentify,
			discord.ScopeEmail,
		),
		twitch.New(
			os.Getenv("TWITCH_CLIENT_ID"),
			os.Getenv("TWITCH_CLIENT_SECRET"),
			callback,
			twitch.ScopeUserReadEmail,
			twitch.ScopeModeratorReadFollowers,
			"user:write:chat",
		),
		spotify.New(
			os.Getenv("SPOTIFY_KEY"),
			os.Getenv("SPOTIFY_SECRET"),
			callback,
			spotify.ScopeUserReadEmail,
			spotify.ScopeUserReadPrivate,
			spotify.ScopeUserReadPlaybackState,
			spotify.ScopeUserModifyPlaybackState,
		),
	)

	server.Handle("GET /login", http.HandlerFunc(handler.Login))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /logout", http.HandlerFunc(handler.Logout))
	return router.NewRouter("/api/oauth2", server), nil
}
