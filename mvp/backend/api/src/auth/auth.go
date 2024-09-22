package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Authenticator interface {
	Provider(http.ResponseWriter) string
	Callback(*http.Request) (*oauth2.Token, error)
	StoreToken(*oauth2.Token) error
}

type oAuth2 struct {
	config *oauth2.Config
}

var oauthStateCookieName string = "oauthstate"

func New(config *oauth2.Config) *oAuth2 {
	return &oAuth2{
		config: config,
	}
}

func (auth *oAuth2) Provider(res http.ResponseWriter) string {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	state := base64.URLEncoding.EncodeToString(bytes)
	cookie := http.Cookie{Name: oauthStateCookieName, Value: state, Expires: expiration}
	http.SetCookie(res, &cookie)

	return auth.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
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

func (auth *oAuth2) StoreToken(token *oauth2.Token) error {
	return errors.New("not implemented, you must implement your own version of StoreToken")
}
