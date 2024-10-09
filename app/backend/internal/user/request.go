package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func GetUserByEmailRequest(accessToken string, email string) (*UserModel, int, error) {
	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), email),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		},
	))
	if err != nil {
		return nil, res.StatusCode, errors.ErrUserNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrUserNotFound
	}

	user, err := decode.Json[UserModel](res.Body)
	if err != nil {
		return &user, res.StatusCode, err
	}
	return &user, res.StatusCode, nil
}

func AddUserRequest(adminTokenRequired bool, user AddUserModel) (*UserModel, int, error) {
	body, err := json.Marshal(user)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to marshal user: %v", err)
	}
	
	res := &http.Response{}
	if adminTokenRequired == false {
		res, err = fetch.Fetch(
			&http.Client{},
			fetch.NewFetchRequest(
				http.MethodPost,
				fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
				bytes.NewReader(body),
				nil,
			),
		)
	}
	if adminTokenRequired == true {
		res, err = fetch.Fetch(
			&http.Client{},
			fetch.NewFetchRequest(
				http.MethodPost,
				fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
				bytes.NewReader(body),
				map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
				},
			),
		)
	}

	if err != nil {
		return nil, res.StatusCode, errors.errCreateuser
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.errCreateuser
	}

	retUser, err := decode.Json[UserModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return &retUser, res.StatusCode, nil
}

func getUserBySessionRequest(accessToken string, sessionId session.SessionModel) (*UserModel, int, error) {
	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/%s", os.Getenv("USER_SERVICE_BASE_URL"), sessionId.UserId),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return nil, res.StatusCode, errors.ErrUserNotFound
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, errors.ErrUserNotFound
	}

	user, err := decode.Json[UserModel](res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return &user, res.StatusCode, nil
}
