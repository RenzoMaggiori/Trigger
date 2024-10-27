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
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/action/action"
	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Watch(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)

	if !ok {
		return errors.ErrAccessTokenCtx
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
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/watch",
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
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
		return errors.ErrGmailWatch
	}

	watchResponse, err := decode.Json[WatchResponse](res.Body)

	if err != nil {
		return err
	}

	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	watchCompleted := workspace.WatchCompletedModel{
		ActionId: actionNode.ActionId,
		UserId:   session.UserId,
		Status:   "active",
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

func fetchUserHistory(accessToken string, lastHistoryId int) (*HistoryList, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/me/history?startHistoryId=%d", lastHistoryId),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
				"Content-Type":  "application/json",
			},
		))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		log.Printf("Watch body: %s", bodyBytes)
		return nil, errors.ErrGmailHistory
	}

	history, err := decode.Json[HistoryList](res.Body)
	if err != nil {
		return nil, errors.ErrGmailHistoryTypeNone
	}
	return &history, nil
}

func getActiveNodes(workspaces []workspace.WorkspaceModel, actionId primitive.ObjectID) []workspace.ActionNodeModel {
	var activeNodes []workspace.ActionNodeModel

	for _, workspace := range workspaces {
		for _, node := range workspace.Nodes {
			if node.ActionId == actionId && node.Status == "active" {
				activeNodes = append(activeNodes, node)
			}
		}
	}
	return activeNodes
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

	workspaces, _, err := workspace.GetByUserId(googleSession.AccessToken, user.Id.Hex())

	if err != nil {
		return err
	}

	activeNodes := getActiveNodes(workspaces, action.Id)

	for _, activeNode := range activeNodes {

		intHistoryId, err := strconv.Atoi(activeNode.Output["historyId"])
		if err != nil {
			return err
		}

		history, err := fetchUserHistory(googleSession.AccessToken, intHistoryId)

		if err != nil {
			return err
		}
		log.Printf("History: %+v", history)

		update := workspace.ActionCompletedModel{
			ActionId: action.Id,
			UserId:   user.Id,
			Output:   map[string]string{"email": ""},
		}

		_, err = workspace.ActionCompletedRequest(googleSession.AccessToken, update)

		if err != nil {
			return err
		}
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
	if res.StatusCode == http.StatusUnauthorized {
		return errors.ErrInvalidGoogleToken
	}
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		log.Printf("Stop body: %s", bodyBytes)
		return errors.ErrGmailStop
	}

	log.Println(accessToken)
	return nil
}