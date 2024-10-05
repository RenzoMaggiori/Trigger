package reaction

import (
	"go.mongodb.org/mongo-driver/mongo"
	"trigger.com/trigger/pkg/action"
)

type GmailReactionCtx string

const AccessTokenCtxKey GmailReactionCtx = GmailReactionCtx("AuthorizationCtxKey")

type Service interface {
	action.Action
}

type Handler struct {
	Service
}

type Model struct {
	Collection *mongo.Collection
}

type WatchBody struct {
	LabelIds  []string `json:"labelIds"`
	TopicName string   `json:"topicName"`
}

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

type HistoryList struct {
	History       []History `json:"history"`
	NextPageToken string    `json:"nextPageToken"`
	HistoryId     string    `json:"historyId"`
}

type History struct {
	Id              string           `json:"id"`
	Messages        []Message        `json:"messages,omitempty"`
	MessagesAdded   []MessageAdded   `json:"messagesAdded,omitempty"`
	MessagesDeleted []MessageDeleted `json:"messagesDeleted,omitempty"`
	LabelsAdded     []LabelAdded     `json:"labelsAdded,omitempty"`
	LabelsRemoved   []LabelRemoved   `json:"labelsRemoved,omitempty"`
}

type Message struct {
	Id           string      `json:"id"`
	ThreadId     string      `json:"threadId"`
	LabelIds     []string    `json:"labelIds"`
	Snippet      string      `json:"snippet"`
	HistoryId    string      `json:"historyId"`
	InternalDate string      `json:"internalDate"`
	Payload      MessagePart `json:"payload"`
	SizeEstimate int64       `json:"sizeEstimate"`
	Raw          string      `json:"raw"`
}

type MessagePart struct {
	PartId   string          `json:"partId"`
	MimeType string          `json:"mimeType"`
	Filename string          `json:"filename"`
	Headers  []MessageHeader `json:"headers"`
	Body     MessagePartBody `json:"body"`
	Parts    []MessagePart   `json:"parts"`
}

type MessageHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessagePartBody struct {
	AttachmentId string `json:"attachmentId"`
	Size         int64  `json:"size"`
	Data         string `json:"data"`
}

type MessageAdded struct {
	Message Message `json:"message"`
}

type MessageDeleted struct {
	Message Message `json:"message"`
}

type LabelAdded struct {
	Message  Message  `json:"message"`
	LabelIds []string `json:"labelIds"`
}

type LabelRemoved struct {
	Message  Message  `json:"message"`
	LabelIds []string `json:"labelIds"`
}
