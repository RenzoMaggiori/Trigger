package reaction

import (
	"trigger.com/trigger/pkg/action"
)

type GmailReactionCtx string

const AccessTokenCtxKey GmailReactionCtx = GmailReactionCtx("AuthorizationCtxKey")

type Service interface {
	action.MultipleReactions
}

type Handler struct {
	Service Service
}

type Model struct {
}

type SendChannelMessageBody struct {
	BroadcasterId string `json:"broadcaster_id"`
	SenderId      string `json:"sender_id"`
	Message       string `json:"message"`
}

type MessageData struct {
	Data []Message `json:"data"`
}

type Message struct {
	MessageID  string  `json:"message_id"`
	IsSent     bool    `json:"is_sent"`
	DropReason *string `json:"drop_reason"` // Use *string to allow for null values
}
