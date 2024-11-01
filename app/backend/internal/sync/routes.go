package sync

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/spotify"
	"github.com/markbates/goth/providers/twitch"
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	syncCollection, ok := ctx.Value(mongodb.CtxKey).(*mongo.Collection)
	if !ok {
		return nil, errors.New("could not find sync mongo collection")
	}

	server := http.NewServeMux()
	middlewares := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{
		Service: Model{
			Collection: syncCollection,
		},
	}

	callback := fmt.Sprintf("http://localhost:%s/api/sync/callback", os.Getenv("SYNC_PORT"))
	CreateProvider(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			callback,
			"https://mail.google.com/",
			"https://www.googleapis.com/auth/documents",
			"https://www.googleapis.com/auth/drive",
			"https://www.googleapis.com/auth/userinfo.profile",
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
			discord.ScopeWebhook,
			discord.ScopeBot,
			discord.ScopeGuilds,
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
		twitch.New(
			os.Getenv("TWITCH_CLIENT_ID"),
			os.Getenv("TWITCH_CLIENT_SECRET"),
			callback,
			twitch.ScopeUserReadEmail,
			twitch.ScopeModeratorReadFollowers,
			"user:write:chat",
		),
	)

	server.Handle("GET /sync-with", http.HandlerFunc(handler.SyncWith))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /{user_id}/{provider}", middlewares(http.HandlerFunc(handler.GetByUserId)))
	return router.NewRouter("/api/sync", server), nil
}
