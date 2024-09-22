package gmail

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/api/src/auth"
	"trigger.com/api/src/service"
)

type Event struct {
	Message struct {
		Data         string `json:"data"`
		MessageId    string `json:"messageId"`
		Message_id   string `json:"message_id"`
		PublishTime  string `json:"publishTime"`
		Publish_time string `json:"publish_time"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type EventData struct {
	EmailAddress string `json:"emailAddress"`
	HistoryId    int64  `json:"historyId"`
}

type Gmail interface {
	auth.Authenticator
	service.Service
	Send() error
}

type Handler struct {
	Gmail
}

type Model struct {
	auth.Authenticator
	Database *mongo.Client
}
