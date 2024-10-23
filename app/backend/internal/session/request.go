package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetSessionByIdRequest(accessToken string, sessionId string) (*SessionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), sessionId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}

	session, err := decode.Json[SessionModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionTypeNone
	}
	return &session, res.StatusCode, nil
}

func GetSessionByAccessTokenRequest(accessToken string) (*SessionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/access_token/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), accessToken),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	session, err := decode.Json[SessionModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionTypeNone
	}
	return &session, res.StatusCode, nil
}

func GetSessionByTokenIdRequest(accessToken string, tokenId string) (*SessionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/token_id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), tokenId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))

	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	session, err := decode.Json[SessionModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionTypeNone
	}
	return &session, res.StatusCode, nil
}

func GetSessionByUserIdRequest(accessToken string, userId string) ([]SessionModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/user_id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), userId),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSessionNotFound
	}

	userSessions, err := decode.Json[[]SessionModel](res.Body)
	if err != nil {
		return userSessions, res.StatusCode, errors.ErrSessionNotFound
	}
	return userSessions, res.StatusCode, nil
}

func UpdateSessionByIdRequest(accessToken string, sessionId string, updateSession UpdateSessionModel) (*SessionModel, int, error) {
	body, err := json.Marshal(updateSession)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := &http.Response{}
	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/session/id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), sessionId),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrFetchingSession
	}

	session, err := decode.Json[SessionModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	return &session, res.StatusCode, nil

}

func AddSessionRequest(accessToken string, addSession AddSessionModel) (*SessionModel, int, error) {
	body, err := json.Marshal(addSession)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrCreatingSession
	}

	res := &http.Response{}
	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/session/add", os.Getenv("SESSION_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrCreatingSession

	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrCreatingSession
	}

	session, err := decode.Json[SessionModel](res.Body)

	if err != nil {
		return nil, res.StatusCode, err
	}

	return &session, res.StatusCode, nil
}
