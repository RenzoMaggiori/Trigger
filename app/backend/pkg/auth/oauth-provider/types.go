package oauthprovider

import (
	"net/http"

	"golang.org/x/oauth2"
	"trigger.com/trigger/pkg/auth/authenticator"
)

type OAuth2Provider interface {
	authenticator.Authenticator
	Callback(*http.Request) (*oauth2.Token, error)
	Config() *oauth2.Config
}

type oAuth2 struct {
	config *oauth2.Config
}

type OAuthProviderCtx string

const OAuthProviderCtxKey OAuthProviderCtx = OAuthProviderCtx("OAuthProviderCtxKey")
