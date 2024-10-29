package gmail

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://mail.google.com/",
			"https://www.googleapis.com/auth/gmail.send",
			"email",
		},
		Endpoint:    google.Endpoint,
		RedirectURL: fmt.Sprintf("http://localhost:%s/api/oauth2/callback", os.Getenv("AUTH_PORT")),
	}
}
