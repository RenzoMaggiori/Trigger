package gmail

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/api/src/auth"
	"trigger.com/api/src/services"
)

type Email struct {
}

type Gmail interface {
	services.Service
	auth.Authenticator
	Webhook() error
	Send(Email) error
}

type Handler struct {
	Gmail
}

type Model struct {
	auth.Authenticator
	database *mongo.Client
}
