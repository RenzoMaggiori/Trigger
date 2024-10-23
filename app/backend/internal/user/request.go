package user

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

func AddUserRequest(accessToken string, addUser AddUserModel) (*UserModel, int, error) {

	body, err := json.Marshal(addUser)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrUserNotFound
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

func GetUserByIdRequest(accessToken string, userId string) (*UserModel, int, error) {
	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/id/%s", os.Getenv("USER_SERVICE_BASE_URL"), userId),
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
