package trigger

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
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
	event, ok := ctx.Value(GmailEventCtxKey).(Event)
	if !ok {
		return errEventCtx
	}
	fmt.Printf("event: %v\n", event)

	data := make([]byte, len(event.Message.Data))
	_, err := base64.NewDecoder(base64.StdEncoding, strings.NewReader(event.Message.Data)).Read(data)
	if err != nil && err != io.EOF {
		return err
	}

	eventData, err := decode.Json[EventData](bytes.NewReader(data))
	if err != nil {
		return err
	}

	fmt.Printf("eventData: %v\n", eventData)

	user, _, err := user.GetUserByEmailRequest(os.Getenv("ADMIN_TOKEN"), eventData.EmailAddress)
	if err != nil {
		return err
	}

	userSessions, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), user.Id.Hex())
	if err != nil {
		return err
	}

	var googleSession *session.SessionModel = nil
	for _, session := range userSessions {
		if *session.ProviderName == "google" {
			googleSession = &session
			break
		}
	}
	if googleSession == nil {
		return errors.ErrSessionNotFound
	}

	action, _, err := action.GetActionByAction(googleSession.AccessToken, googleTriggerActionName)

	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		UserId:   user.Id,
		Output:   map[string]any{"hello": "world"},
	}

	_, err = workspace.ActionCompletedRequest(googleSession.AccessToken, update)

	if err != nil {
		return err
	}

	return nil
}

func (m Model) Stop(ctx context.Context) error {
	accessToken := ctx.Value(AccessTokenCtxKey)

	fmt.Printf("Access token: %s", accessToken)
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/stop",
			nil,
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
		log.Printf("Stop body: %s", bodyBytes)
		return errGmailStop
	}

	log.Println(accessToken)
	return nil
}