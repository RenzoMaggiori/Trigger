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
	"os"
	"strings"

	"golang.org/x/oauth2"
	"trigger.com/api/src/endpoints/user"
	"trigger.com/api/src/lib"
)

var (
	_                   Service = Model{}
	gmailAccessTokenKey string  = "gmailAccessTokenKey"
	gmailEventKey       string  = "gmailEventKey"
)

func (m Model) GetUserFromGoogle(token *oauth2.Token) (*GmailUser, error) {
	res, err := lib.Fetch(lib.NewFetchRequest(
		http.MethodGet,
		"https://gmail.googleapis.com/gmail/v1/users/me/profile",
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
		},
	))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	user, err := lib.JsonDecode[GmailUser](res.Body)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m Model) GetUserFromDbByEmail(email string) (*user.User, error) {
	res, err := lib.Fetch(lib.NewFetchRequest(
		"GET",
		fmt.Sprintf("%s/user/%s", os.Getenv("API_URL"), email),
		nil,
		nil,
	))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code, received %s\n", res.Status)
	}

	user, err := lib.JsonDecode[user.User](res.Body)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m Model) AddUserToDb(email string, token *oauth2.Token) error {
	res, err := lib.Fetch(lib.NewFetchRequest(
		"POST",
		fmt.Sprintf("%s/user", os.Getenv("API_URL")),
		map[string]any{
			"email":        email,
			"accessToken":  token.AccessToken,
			"refreshToken": token.RefreshToken,
			"tokenType":    token.TokenType,
			"expiry":       token.Expiry,
		},
		nil,
	))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code, received %s\n", res.Status)
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

	user, err := m.GetUserFromDbByEmail(eventData.EmailAddress)
	if err != nil {
		return err
	}

	token := oauth2.Token{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
		TokenType:    user.TokenType,
		Expiry:       user.Expiry,
	}
	client := m.Authenticator.Config().Client(context.TODO(), &token)

	newEmail, err := fetchUserHistory(user.AccessToken, eventData)
	if err != nil {
		return err
	}
	if !newEmail {
		return errors.New("no new emails on inbox")
	}
	// TODO: call send
	return nil
}

func (m Model) Send() error {
	return errors.New("Not implemented")
}
