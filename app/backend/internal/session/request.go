package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
	"github.com/markbates/goth"
	"trigger.com/trigger/internal/user"
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

//? Is correct the return value?
func GetSessionByIdRequest(accessToken string, sessionId string, updateSession UpdateSessionModel) (*UpdateSessionModel, int, error) {
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
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		log.Println("Credentials Login fetch [:8082/api/session/id/{id}] error")
		return nil, http.StatusInternalServerError, errors.ErrSessionNotRetrieved
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrSessionNotRetrieved
	}

	session, err := decode.Json[UpdateSessionModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, errors.ErrSessionTypeNone
	}
	return &session, res.StatusCode, nil

}

//? Is correct the return value?
func addSessionRequest(accessToken string, addSession AddSessionModel, gothUser goth.User) (string, int, error) {
	body, err := json.Marshal(addSession)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrSessionNotCreated
	}

	res := &http.Response{}
	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/session/add", os.Getenv("SESSION_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrSessionNotCreated

	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errors.New("unable to create session")
	}

	return gothUser.AccessToken, res.StatusCode, nil
}
