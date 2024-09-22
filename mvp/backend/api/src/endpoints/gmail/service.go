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
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"trigger.com/api/src/lib"
)

var (
	_                   Gmail  = Model{}
	gmailAccessTokenKey string = "gmailAccessTokenKey"
	gmailEventKey       string = "gmailEventKey"
)

func addToDB(userCollection *mongo.Collection, user GmailUser, token *oauth2.Token) error {
	var existingUser bson.M
	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.EmailAddress}}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		newUser := bson.D{
			{Key: "email", Value: user.EmailAddress},
			{Key: "access_token", Value: token.AccessToken},
			{Key: "refresh_token", Value: token.RefreshToken},
			{Key: "token_expiry", Value: token.Expiry},
			{Key: "token_type", Value: token.TokenType},
		}

		_, err := userCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			return err
		}
	} else if err == nil {
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "access_token", Value: token.AccessToken},
				{Key: "refresh_token", Value: token.RefreshToken},
				{Key: "token_expiry", Value: token.Expiry},
				{Key: "token_type", Value: token.TokenType},
			}},
		}

		_, err := userCollection.UpdateOne(context.TODO(), bson.D{{Key: "email", Value: user.EmailAddress}}, update)
		if err != nil {
			return err
		}
	}
	return nil
}

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
	if err != nil {
		return err
	}
	userCollection := m.Mongo.Database(os.Getenv("MONGO_DB")).Collection("user")
	if err := addToDB(userCollection, user, token); err != nil {
		return err
	}
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
		return fmt.Errorf("invalid response, got %s", res.Status)
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
