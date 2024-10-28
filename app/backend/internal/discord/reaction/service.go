package reaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error {
	log.Println("actionName: ", actionName)
	switch actionName {
	case "send_message":
		return m.sendMessage(ctx, action)
	}

	return nil
}

func (m Model) sendMessage(ctx context.Context, actionNode workspace.ActionNodeModel) error {
	content := actionNode.Input["content"]
	channel_id := actionNode.Input["channel_id"]
	ttsStr := actionNode.Input["tts"]
    tts, err := strconv.ParseBool(ttsStr)

	if err != nil {
		return err
	}

	payload := MessagegContent{
		Content: content,
		TTS: tts,
	}
    body, err := json.Marshal(payload)
    if err != nil {
        return err
    }

	return manageNewMessage(channel_id, body)
}

func manageNewMessage(channelID string, body []byte) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/channels/%s/messages", baseURL, channelID),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": "Bot " + os.Getenv("BOT_TOKEN"),
				"Content-Type": "application/json",
			},
		),
	)

	if err != nil {
		return err
	}

    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send message, status: %d", res.StatusCode)
    }

    fmt.Println("Message sent successfully!")
    return nil
}


func userInfo(accessToken string) (UserInfo, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			userEndpoint,
			nil,
			map[string]string{
				"Authorization": "Bearer " + accessToken,
			},
		),
	)

	if err != nil {
		return UserInfo{}, err
	}

	defer res.Body.Close()

	userInfo, err := decode.Json[UserInfo](res.Body)

	if err != nil {
		return UserInfo{}, err
	}

	return userInfo, nil
}

func guilds(accessToken string) (string, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/guilds", userEndpoint),
			nil,
			map[string]string{
				"Authorization": "Bearer " + accessToken,
			},
		),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func createWebhook(accessToken string, channelId string, webhookName string) error {
	createWebhook := NewWebhook{
		Name: webhookName,
		Avatar: "",
	}

	body, err := json.Marshal(createWebhook)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/channels/%s/webhooks", baseURL, channelId),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": "Bearer " + accessToken,
				"Content-Type":  "application/json",
			},
		),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)

	log.Println("create_webhook: string(body)")
	log.Println(string(body))

	if err != nil {
		return err
	}

	return nil
}

func deleteWebhook(accessToken string, webhookId string, webhookToken string) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/webhooks/%s/%s", baseURL, webhookId, webhookToken),
			nil,
			map[string]string{
				"Authorization": "Bearer " + accessToken,
			},
		),
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (m Model) RefreshToken(accessToken string, webhookId string, webhookToken string) (Webhook, error) {
	// API_ENDPOINT := "https://discord.com/api/v10"
	// CLIENT_ID := "332269999912132097"
	// CLIENT_SECRET := "937it3ow87i4ery69876wqire"
	// REDIRECT_URI := "https://localhost:3000"

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			tokenURL,
			nil,
			map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
		),
	)

	if err != nil {
		return Webhook{}, err
	}

	defer res.Body.Close()

	webhook, err := decode.Json[Webhook](res.Body)

	if err != nil {
		return Webhook{}, err
	}

	return webhook, nil
}
