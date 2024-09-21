package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Authenticator interface {
	Provider(res http.ResponseWriter) string
	Callback(ctx context.Context, authCode string) (*oauth2.Token, error)
}

type oAuth2 struct {
	config *oauth2.Config
}

func New(config *oauth2.Config) *oAuth2 {
	return &oAuth2{
		config: config,
	}
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (auth *oAuth2) Provider(res http.ResponseWriter) string {
	oauthState := generateStateOauthCookie(res)
	return auth.config.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
}

func (auth *oAuth2) Callback(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := auth.config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}
	return token, nil
}
