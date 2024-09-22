package gmail

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"trigger.com/api/src/auth"
	"trigger.com/api/src/database"
	"trigger.com/api/src/middleware"
)

func AuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://mail.google.com/",
			"https://www.googleapis.com/auth/gmail.send",
			"https://www.googleapis.com/auth/gmail.modify",
		},
		Endpoint:    google.Endpoint,
		RedirectURL: fmt.Sprintf("%s/api/auth/gmail/callback", os.Getenv("API_URL")),
	}

}

func Router(ctx context.Context) (*http.ServeMux, error) {
	database, ok := ctx.Value(database.CtxKey).(*mongo.Client)
	if !ok {
		return nil, fmt.Errorf("could not get Database from Context")
	}

	router := http.NewServeMux()
	authMiddleware := middleware.Create(
		middleware.Auth,
	)
	handler := Handler{Gmail: Model{
		Authenticator: auth.New(AuthConfig()),
		Database:      database,
	}}

	router.HandleFunc("GET /auth/gmail/provider", handler.AuthProvider)
	router.HandleFunc("GET /auth/gmail/callback", handler.AuthCallback)
	router.Handle("GET /gmail/register", authMiddleware(http.HandlerFunc(handler.Register)))
	return router, nil
}
