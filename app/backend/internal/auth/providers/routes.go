package providers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitch"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()

	handler := Handler{
		Service: Model{},
	}

	callback := fmt.Sprintf("http://localhost:%s/api/oauth2/callback", os.Getenv("AUTH_PORT"))
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
	)

	server.Handle("GET /login", http.HandlerFunc(handler.Login))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /logout", http.HandlerFunc(handler.Logout))
	return router.NewRouter("/oauth2", server), nil
}
