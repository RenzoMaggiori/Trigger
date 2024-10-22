package oauthprovider

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const oauthStateCookieName string = "oauthstate"

var (
	errResponseNotFound error = errors.New("could not find response in context")
)

func New(config *oauth2.Config) *oAuth2 {
	return &oAuth2{
		config: config,
	}
}

func (auth *oAuth2) Login(ctx context.Context) (string, error) {
	w, ok := ctx.Value(OAuthProviderCtxKey).(http.ResponseWriter)
	if !ok {
		return "", errResponseNotFound
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	state := base64.URLEncoding.EncodeToString(bytes)
	cookie := http.Cookie{Name: oauthStateCookieName, Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return auth.config.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

func (auth *oAuth2) Callback(req *http.Request) (*oauth2.Token, error) {
	oauthState, err := req.Cookie(oauthStateCookieName)
	if err != nil {
		return nil, err
	}

	state := req.FormValue("state")
	code := req.FormValue("code")
	if state != oauthState.Value {
		return nil, errors.New("invalid oauth google state")
	}

	token, err := auth.config.Exchange(req.Context(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (auth *oAuth2) Config() *oauth2.Config {
	return auth.config
}
