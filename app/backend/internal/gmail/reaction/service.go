package reaction

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

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

func (m Model) Action(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	accessToken := ctx.Value(AccessTokenCtxKey).(string)

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/session/access_token/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), accessToken),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errSessionNotFound
	}
	session, err := decode.Json[session.SessionModel](res.Body)

	if err != nil {
		return err
	}

	res, err = fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/%s", os.Getenv("USER_SERVICE_BASE_URL"), session.UserId),
			nil,
			map[string]string{
				"Authorization": accessToken,
			},
		),
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errUserNotFound
	}
	user, err := decode.Json[user.UserModel](res.Body)

	if err != nil {
		return err
	}

	rawEmail, err := createRawEmail(user.Email, user.Email, "Hello world", "AAAAAAAAAA")

	if err != nil {
		return fmt.Errorf("failed to create raw email: %v", err)
	}
	requestBody := fmt.Sprintf(`{"raw": "%s"}`, rawEmail)

	body, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	res, err = fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/messages/send",
			bytes.NewReader(body),
			map[string]string{
				"Authorization": accessToken,
				"Content-Type":  "application/json",
			},
		),
	)
	if err != nil {
		return err
	}

	return nil
}
