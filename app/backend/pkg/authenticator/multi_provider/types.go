package multiprovider

import (
	"trigger.com/trigger/pkg/authenticator"
)

type Service interface {
	authenticator.Authenticator
}

type Handler struct {
	Service
}

type Model struct {
}
