package session

import (
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetSessionByTokenRequest(accessToken string) (*SessionModel, int, error) {
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
