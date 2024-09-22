package gmail

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2"
	"trigger.com/api/src/lib"
)

var (
	_                   Gmail  = Model{}
	gmailAccessTokenKey string = "gmailAccessTokenKey"
	gmailEventKey       string = "gmailEventKey"
)

func (m Model) StoreToken(token *oauth2.Token) error {
	res, err := lib.Fetch(lib.NewFetchRequest(
		http.MethodGet,
		"https://gmail.googleapis.com/gmail/v1/users/me/profile",
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
		},
	))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	user, err := lib.JsonDecode[GmailUser](res.Body)
	userCollection := m.Mongo.Database(os.Getenv("MONGO_DB")).Collection("user")

	result := userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.EmailAddress}})

	// store the token info and user in db
	return nil
}

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

func (m Model) Webhook(ctx context.Context) error {
	event, ok := ctx.Value(gmailEventKey).(Event)
	if !ok {
		return errors.New("could not retrieve event")
	}
	fmt.Printf("event: %v\n", event)

	// TODO: get access_token from db
	// TODO: check access_token is valid else use refresh token

	data := make([]byte, len(event.Message.Data))
	_, err := base64.NewDecoder(base64.StdEncoding, strings.NewReader(event.Message.Data)).Read(data)
	if err != nil && err != io.EOF {
		return err
	}

	eventData, err := lib.JsonDecode[EventData](bytes.NewReader(data))
	if err != nil {
		return err
	}
	fmt.Printf("eventData: %v\n", eventData)

	// TODO: fetch history

	return nil
}

func (m Model) Send() error {
	return errors.New("Not implemented")
}
