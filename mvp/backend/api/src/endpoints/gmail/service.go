package gmail

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"trigger.com/api/src/lib"
)

var _ Gmail = Model{}

var gmailAccessTokenKey string = "gmailAccessTokenKey"

func (m Model) Register(ctx context.Context) error {
	accessToken, ok := ctx.Value(gmailAccessTokenKey).(string)
	if !ok {
		return errors.New("could not retrieve access token")
	}

	res, err := lib.Fetch(lib.NewFetchRequest(
		http.MethodPost,
		"https://gmail.googleapis.com/gmail/v1/users/me/watch",
		map[string]any{
			"labelIds":  []string{"INBOX"},
			"topicName": "projects/trigger-436310/topics/Trigger",
		},
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			"Content-Type":  "application/json",
		}))
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
