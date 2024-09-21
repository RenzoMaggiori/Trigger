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
