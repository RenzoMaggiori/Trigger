package providers

import (
	"context"
	"net/http"
	"os"

	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"trigger.com/trigger/pkg/router"
)

func Router(ctx context.Context) (*router.Router, error) {
	server := http.NewServeMux()

	handler := Handler{
		Service: Model{},
	}

	CreateProvider(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:8000/api/oauth2/callback",
			// "https://mail.google.com/",
			"https://www.googleapis.com/auth/gmail.send",
			// "email",
		),
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			"http://localhost:8000/api/oauth2/callback",
		),
	)

	server.Handle("GET /login", http.HandlerFunc(handler.Login))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /logout", http.HandlerFunc(handler.Logout))
	return router.NewRouter("/oauth2", server), nil
}
