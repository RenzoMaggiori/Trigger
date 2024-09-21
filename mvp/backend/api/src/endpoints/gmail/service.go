package gmail

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var _ Gmail = Model{}


/* "watch": {
"id": "gmail.users.watch",
"path": "gmail/v1/users/{userId}/watch",
"flatPath": "gmail/v1/users/{userId}/watch",
"httpMethod": "POST",
"parameters": {
"userId": {
"description": "The user's email address. The special value `me` can be used to indicate the authenticated user.",
"default": "me",
"location": "path",
"required": true,
"type": "string"
}
},
"parameterOrder": [
"userId"
],
"request": {
"$ref": "WatchRequest"
},
"response": {
"$ref": "WatchResponse"
},
"scopes": [
"https://mail.google.com/",
"https://www.googleapis.com/auth/gmail.metadata",
"https://www.googleapis.com/auth/gmail.modify",
"https://www.googleapis.com/auth/gmail.readonly"
],
"description": "Set up or update a push notification watch on the given user mailbox."
}, */

/* "WatchRequest": {
"id": "WatchRequest",
"description": "Set up or update a new push notification watch on this user's mailbox.",
"type": "object",
"properties": {
"labelIds": {3 items},
"labelFilterAction": {5 items},
"labelFilterBehavior": {4 items},
"topicName": {2 items}
}
}, */

/* "WatchResponse": {
"id": "WatchResponse",
"description": "Push notification watch response.",
"type": "object",
"properties": {
"historyId": {
"description": "The ID of the mailbox's current history record.",
"type": "string",
"format": "uint64"
},
"expiration": {
"description": "When Gmail will stop sending notifications for mailbox updates (epoch millis). Call `watch` again before this time to renew the watch.",
"type": "string",
"format": "int64"
}
}
}, */

func (m Model) Register(ctx context.Context) error {
	return errors.New("Not implemented")
}

func (m Model) Webhook() error {
	return errors.New("Not implemented")
}

func (m Model) Send(email Email) error {
	return errors.New("Not implemented")
}
