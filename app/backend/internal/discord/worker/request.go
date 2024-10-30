package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetCurrDiscordSessionReq(accessToken string) (*DiscordSessionModel, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/current", workerBaseURL),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.ErrDiscordUserSessionNotFound
	}

	session, err := decode.Json[DiscordSessionModel](res.Body)
	if err != nil {
		return nil, errors.ErrDecodeData
	}

	return &session, nil
}

func AddDiscordSessionReq(accessToken string, session AddDiscordSessionModel) error {
	body, err := json.Marshal(session)
	if err != nil {
		return errors.ErrMarshalData
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/add", workerBaseURL),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return errors.ErrAddDiscordSession
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.ErrAddDiscordSession
	}

	return nil
}

func UpdateDiscordSessionReq(accessToken string, id string, session UpdateDiscordSessionModel) error {
	body, err := json.Marshal(session)
	if err != nil {
		return errors.ErrMarshalData
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPut,
			fmt.Sprintf("%s/update/%s", workerBaseURL, id),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return errors.ErrUpdateDiscordSession
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.ErrUpdateDiscordSession
	}

	return nil
}

func DeleteDiscordSessionReq(accessToken string, id string) error {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/delete/%s", workerBaseURL, id),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return errors.ErrDeleteDiscordSession
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.ErrDeleteDiscordSession
	}

	return nil
}

func GetMeReq(token string) (*Me, error) {
	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/me", workerBaseURL),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", token),
			},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", errors.ErrDiscordMe, res.StatusCode)
	}

	me, err := decode.Json[Me](res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDecodeData, err)
	}

	return &me, nil
}
