package gmail

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

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

func refreshAccessToken(refreshToken string) (*oauth2.Token, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	tokenURL := "https://oauth2.googleapis.com/token"

	values := url.Values{}
	values.Set("client_id", clientID)
	values.Set("client_secret", clientSecret)
	values.Set("refresh_token", refreshToken)
	values.Set("grant_type", "refresh_token")

	resp, err := http.PostForm(tokenURL, values)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %v", err)
	}
	defer resp.Body.Close()

	var newToken oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&newToken); err != nil {
		return nil, fmt.Errorf("failed to decode new token: %v", err)
	}

	return &newToken, nil
}

func fetchUserHistory(accessToken string, eventData EventData) (bool, error) {
	url := fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/%s/history?startHistoryId=%d", eventData.EmailAddress, eventData.HistoryId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to fetch Gmail history: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to fetch Gmail history: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch Gmail history, got status: %s", resp.Status)
	}

	var historyResponse HistoryList
	if err := json.NewDecoder(resp.Body).Decode(&historyResponse); err != nil {
		return false, fmt.Errorf("failed to decode Gmail history response: %v", err)
	}

	// * Here we check if the history list we got has at the start an Added message (new email received)
	if len(historyResponse.History) > 0 {
		firstHistoryItem := historyResponse.History[0]

		if len(firstHistoryItem.MessagesAdded) > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

func (m Model) Webhook(ctx context.Context) error {
	event, ok := ctx.Value(gmailEventKey).(Event)
	if !ok {
		return errors.New("could not retrieve event")
	}
	fmt.Printf("event: %v\n", event)

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

	userCollection := m.Mongo.Database(os.Getenv("MONGO_DB")).Collection("user")
	var user DbUser

	err = userCollection.FindOne(ctx, bson.D{{Key: "email", Value: eventData.EmailAddress}}).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to retrieve user from DB: %v", err)
	}
	// * here we check if the token is valid
	if time.Now().After(user.Expiry) {
		newToken, err := refreshAccessToken(user.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to refresh access token: %v", err)
		}
		// * after the refresh we update de DB with the new token
		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "access_token", Value: newToken.AccessToken},
				{Key: "expiry", Value: newToken.Expiry},
			}}, // ! Maybe add new refresh token if needed?
		}
		_, err = userCollection.UpdateOne(ctx, bson.D{{Key: "email", Value: eventData.EmailAddress}}, update)
		if err != nil {
			return fmt.Errorf("failed to update access token in DB: %v", err)
		}
		user.AccessToken = newToken.AccessToken
	}
	newEmail, err := fetchUserHistory(user.AccessToken, eventData)
	if err != nil {
		return err
	}
	if !newEmail {
		return errors.New("no new emails on inbox")
	}
	return nil
}

func (m Model) Send() error {
	return errors.New("Not implemented")
}
