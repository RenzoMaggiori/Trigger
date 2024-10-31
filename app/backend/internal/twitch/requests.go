package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetUserByAccessTokenRequest(syncModel sync.SyncModel) (*UserData, error) {
	client, err := oaclient.New(context.TODO(), Config(), &syncModel)
	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodGet,
			"https://api.twitch.tv/helix/users",
			nil,
			map[string]string{
				"Client-Id":    os.Getenv("TWITCH_CLIENT_ID"),
				"Content-Type": "application/json",
			},
		),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.ErrTwitchUser
	}

	userResponse, err := decode.Json[UserData](res.Body)
	if err != nil {
		return nil, err
	}

	if len(userResponse.Data) == 0 {
		return nil, errors.ErrTwitchUser
	}
	return &userResponse, nil
}

func GetAppAccessTokenRequest() (*AppAccessTokenResponse, error) {
	appAccessTokenBody := AppAccessTokenBody{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		GrantType:    "client_credentials",
	}
	body, err := json.Marshal(appAccessTokenBody)
	if err != nil {
		return nil, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://id.twitch.tv/oauth2/token",
			bytes.NewReader(body),
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)

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
		return nil, errors.ErrTwitchAppAccessToken
	}

	appAccessTokenResponse, err := decode.Json[AppAccessTokenResponse](res.Body)
	if err != nil {
		return nil, err
	}
	return &appAccessTokenResponse, nil
}
