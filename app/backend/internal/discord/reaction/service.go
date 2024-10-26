package reaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) UserInfo(accessToken string) (UserInfo, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			"https://discord.com/api/users/@me",
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

func (m Model) Guilds(accessToken string) (string, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			"https://discord.com/api/users/@me/guilds",
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

func (m Model) CreateWebhook(accessToken string, channelId string, webhookName string) (string, error) {
	createWebhook := map[string]interface{}{
		"name":   webhookName,
		"avatar": nil,
	}

	body, err := json.Marshal(createWebhook)
	if err != nil {
		return "", err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("https://discord.com/api/channels/%s/webhooks", channelId),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": "Bearer " + accessToken,
				"Content-Type":  "application/json",
			},
		),
	)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (m Model) DeleteWebhook(accessToken string, webhookId string, webhookToken string) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", webhookId, webhookToken),
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

func (m Model) SendMessage(accessToken string, webhookId string, webhookToken string, content string, username string) error {
	message := map[string]interface{}{
		"content": content,
		"username": username,
		"avatar_url": nil,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", webhookId, webhookToken),
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

	return nil
}

func (m Model) RefreshToken(accessToken string, webhookId string, webhookToken string) (Webhook, error) {
	// API_ENDPOINT := "https://discord.com/api/v10"
	// CLIENT_ID := "332269999912132097"
	// CLIENT_SECRET := "937it3ow87i4ery69876wqire"
	// REDIRECT_URI := "https://nicememe.website"

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			"https://discord.com/api/oauth2/token",
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