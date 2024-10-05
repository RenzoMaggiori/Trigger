package trigger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken := ctx.Value(AccessTokenCtxKey)

	watchBody := WatchBody{
		LabelIds:  []string{"INBOX"},
		TopicName: "projects/trigger-436310/topics/Trigger",
	}

	body, err := json.Marshal(watchBody)

	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/watch",
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("%s", accessToken),
				"Content-Type":  "application/json",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		log.Printf("Watch body: %s", bodyBytes)
		return errGmailWatch
	}

	log.Println(accessToken)
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	// TODO: Handle webhook
	return nil
}
