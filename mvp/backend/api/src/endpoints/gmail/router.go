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
)

var googleAuthConfig = &oauth2.Config{
	Scopes: []string{
		"https://mail.google.com/",
		"https://www.googleapis.com/auth/gmail.send",
		"https://www.googleapis.com/auth/gmail.modify",
	},
	Endpoint:    google.Endpoint,
	RedirectURL: "http://localhost:8000/api/auth/google/callback",
}

func Router(ctx context.Context) (*http.ServeMux, error) {
	googleAuthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleAuthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	database, ok := ctx.Value(database.CtxKey).(*mongo.Client)
	if !ok {
		return nil, fmt.Errorf("could not get Database from Context")
	}

	router := http.NewServeMux()
	handler := Handler{Gmail: Model{
		Authenticator: auth.New(googleAuthConfig),
		database:      database,
	}}

	router.HandleFunc("GET /auth/gmail/provider", handler.AuthProvider)
	router.HandleFunc("GET /auth/gmail/callback", handler.AuthCallback)
	return router, nil
}
