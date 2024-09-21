package gmail

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

var gmailAccessTokenKey string = "gmailAccessTokenKey"

func (m Model) Register(ctx context.Context) error {
	accessToken, ok := ctx.Value(gmailAccessTokenKey).(string)
	if !ok {
		return errors.New("could not retrieve access token")
	}

	body, err := json.Marshal(map[string]any{"labelIds": []string{"INBOX"}, "topicName": "projects/trigger-436310/topics/Trigger"})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://gmail.googleapis.com/gmail/v1/users/me/watch", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response, got %s\n", res.Status)
	}
	return nil
}

func (m Model) Webhook() error {
	return errors.New("Not implemented")
}

func (m Model) Send(email Email) error {
	return errors.New("Not implemented")
}
