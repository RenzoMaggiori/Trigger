package reaction

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"

	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)

	if !ok {
		return errors.ErrAccessTokenCtx
	}

	switch actionName {
	case "send_email":
		return m.SendGmail(ctx, accessToken, action)
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

	// Use URL encoding to avoid manual replacements of + and /
	rawMessage := base64.URLEncoding.EncodeToString(email.Bytes())

	return rawMessage, nil
}

func (m Model) SendGmail(ctx context.Context, accessToken string, actionNode workspace.ActionNodeModel) error {
	session, _, err := session.GetSessionByAccessTokenRequest(accessToken)

	if err != nil {
		return err
	}

	user, _, err := user.GetUserByIdRequest(accessToken, session.UserId.Hex())

	if err != nil {
		return err
	}

	// TODO: Replace hardcoded values with: actionNode.Inputs["from"], actionNode.Inputs["to"], ...
	rawEmail, err := createRawEmail(user.Email,
		"johndoe@gmail.com", "asd", "asd")

	if err != nil {
		return errors.ErrCreatingEmail
	}

	requestBody := map[string]string{
		"raw": rawEmail,
	}

	body, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/messages/send",
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
		log.Printf("Send gmail body: %s", bodyBytes)
		return errors.ErrFailedToSendEmail
	}

	return nil
}
