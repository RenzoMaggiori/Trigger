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

	"golang.org/x/oauth2"
	"trigger.com/api/src/endpoints/user"
	"trigger.com/api/src/lib"
)

/**
* ! If you are reading this, our deepest sympathies are with you.
* ! What lies below is not code; it is a monument to our desperation, a relic of our darkest hour.
* ! We sought to connect email with email, but in doing so, we lost ourselves in an abyss of complexity.
* ! To the brave soul tasked with maintaining this, may you find the strength to endure where we could not.
* ! Abandon all hope, ye who enter here.
**/

var (
	_                   Service = Model{}
	gmailAccessTokenKey string  = "gmailAccessTokenKey"
	gmailEventKey       string  = "gmailEventKey"
)

func (m Model) GetUserFromGoogle(token *oauth2.Token) (*GmailUser, error) {
	res, err := lib.Fetch(
		&http.Client{},
		lib.NewFetchRequest(
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
	res, err := lib.Fetch(
		&http.Client{},
		lib.NewFetchRequest(
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
		return nil, fmt.Errorf("invalid status code, received %s", res.Status)
	}

	user, err := lib.JsonDecode[user.User](res.Body)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m Model) AddUserToDb(email string, token *oauth2.Token) error {
	res, err := lib.Fetch(
		&http.Client{},
		lib.NewFetchRequest(
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
		return fmt.Errorf("invalid status code, received %s", res.Status)
	}
	return nil
}

func (m Model) Register(ctx context.Context) error {
	accessToken, ok := ctx.Value(gmailAccessTokenKey).(string)
	if !ok {
		return errors.New("could not retrieve access token")
	}

	res, err := lib.Fetch(
		&http.Client{},
		lib.NewFetchRequest(
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

func fetchUserHistory(eventData EventData, client *http.Client) (bool, error) {
	res, err := lib.Fetch(
		client,
		lib.NewFetchRequest(
			"GET",
			fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/me/history?startHistoryId=%d", eventData.HistoryId),
			nil,
			nil,
		))
	if err != nil {
		return false, fmt.Errorf("failed to fetch Gmail history: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to fetch Gmail history, got status: %s", res.Status)
	}

	history, err := lib.JsonDecode[HistoryList](res.Body)
	if err != nil {
		return false, fmt.Errorf("failed to decode Gmail history response: %v", err)
	}

	// * Here we check if the history list we got has at the start an Added message (new email received)
	if len(history.History) > 0 {
		firstHistoryItem := history.History[0]

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

	newEmail, err := fetchUserHistory(eventData, client)
	if err != nil {
		return err
	}
	// * If we have a new email we send an email as a response
	if newEmail {
		m.Send(client, eventData.EmailAddress, eventData.EmailAddress, "New Message", "You received a new email, go check it")
	}
	return nil
}

func createRawEmail(from string, to string, subject string, body string) (string, error) {
	var email bytes.Buffer
	email.WriteString(fmt.Sprintf("From: %s\r\n", from))
	email.WriteString(fmt.Sprintf("To: %s\r\n", to))
	email.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	email.WriteString("\r\n")
	email.WriteString(body)

	rawMessage := base64.StdEncoding.EncodeToString(email.Bytes())

	// * Gmail's API requires the base64-encoded message to be in a URL-safe format without padding
	// * So we replace this characters with safe ones for URL-safe base64 encoding
	rawMessage = strings.ReplaceAll(rawMessage, "+", "-")
	rawMessage = strings.ReplaceAll(rawMessage, "/", "_")
	rawMessage = strings.TrimRight(rawMessage, "=")

	return rawMessage, nil
}

// * Here we just create an email and send it to the user
// ? Can you send an email through the user itself? (probably not)
// TODO: test that all works as intended
func (m Model) Send(client *http.Client, from string, to string, subject string, body string) error {
	rawEmail, err := createRawEmail(from, to, subject, body)
	if err != nil {
		return fmt.Errorf("failed to create raw email: %v", err)
	}

	url := "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"
	requestBody := fmt.Sprintf(`{"raw": "%s"}`, rawEmail)

	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send email, got status: %s", resp.Status)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
