package github

import (
	oauthprovider "trigger.com/trigger/pkg/auth/oauth-provider"
)

type Service interface {
	oauthprovider.OAuth2Provider
}

type Handler struct {
	Service
}

type Model struct {
	oauthprovider.OAuth2Provider
}
