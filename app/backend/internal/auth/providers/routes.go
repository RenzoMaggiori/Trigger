package providers

import (
	"context"
	"fmt"
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

	callback := fmt.Sprintf("%s/api/oauth2/callback", os.Getenv("AUTH_SERVICE_BASE_URL"))
	CreateProvider(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			callback,
		),
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callback,
		),
	)

	server.Handle("GET /login", http.HandlerFunc(handler.Login))
	server.Handle("GET /callback", http.HandlerFunc(handler.Callback))
	server.Handle("GET /logout", http.HandlerFunc(handler.Logout))
	return router.NewRouter("/oauth2", server), nil
}
