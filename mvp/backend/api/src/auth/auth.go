package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type Authenticator interface {
	Provider() string
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

func (auth *oAuth2) Provider() string {
	return auth.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (auth *oAuth2) Callback(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := auth.config.Exchange(ctx, authCode)
	if err != nil {
		return nil, err
	}
	return token, nil
}
