package gmail

import "trigger.com/api/src/services"

type Email struct {
}

type Gmail interface {
	services.Service
	Webhook() error
	Send(Email) error
}

type Handler struct {
	Gmail
}

type Model struct {
}
