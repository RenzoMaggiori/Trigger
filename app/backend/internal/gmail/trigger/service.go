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
	"trigger.com/trigger/internal/gmail"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "google")
	if err != nil {
		return err
	}

	client, err := oaclient.New(ctx, gmail.Config(), syncModel)
	if err != nil {
		return err
	}

	watchBody := WatchBody{
		LabelIds:  []string{"INBOX"},
		TopicName: "projects/trigger-436310/topics/Trigger",
	}
	body, err := json.Marshal(watchBody)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/watch",
			bytes.NewReader(body),
			map[string]string{
				"Content-Type": "application/json",
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
		return errors.ErrGmailWatch
	}

	watchResponse, err := decode.Json[WatchResponse](res.Body)
	if err != nil {
		return err
	}

	watchCompleted := workspace.WatchCompletedModel{
		ActionId: actionNode.ActionId,
		UserId:   session.UserId,
		Output: map[string]string{
			"historyId":  watchResponse.HistoryId,
			"expiration": watchResponse.Expiration,
		},
	}
	_, _, err = workspace.WatchCompletedRequest(accessToken, watchCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Webhook(ctx context.Context) error {
	event, ok := ctx.Value(GmailEventCtxKey).(Event)
	if !ok {
		return errors.ErrEventCtx
	}

	data := make([]byte, len(event.Message.Data))
	_, err := base64.NewDecoder(base64.StdEncoding, strings.NewReader(event.Message.Data)).Read(data)
	if err != nil && err != io.EOF {
		return err
	}

	eventData, err := decode.Json[EventData](bytes.NewReader(data))
	if err != nil {
		return err
	}

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

	action, _, err := action.GetByActionNameRequest(googleSession.AccessToken, googleWatchActionName)
	if err != nil {
		return err
	}

	update := workspace.ActionCompletedModel{
		ActionId: action.Id,
		UserId:   user.Id,
		Output:   map[string]string{"email": ""},
	}
	_, err = workspace.ActionCompletedRequest(googleSession.AccessToken, update)
	if err != nil {
		return err
	}
	return nil
}

func (m Model) Stop(ctx context.Context) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, session.UserId.Hex(), "google")
	if err != nil {
		return err
	}

	client, err := oaclient.New(ctx, gmail.Config(), syncModel)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/stop",
			nil,
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusUnauthorized {
		return errors.ErrInvalidGoogleToken
	}
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v: %s", errors.ErrGmailStop, bodyBytes)
	}
	return nil
}
