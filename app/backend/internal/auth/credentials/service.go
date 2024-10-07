package credentials

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/hash"
	"trigger.com/trigger/pkg/jwt"
)

func (m Model) Login(ctx context.Context) (string, error) {
	credentials, ok := ctx.Value(CredentialsCtxKey).(CredentialsModel)
	if !ok {
		return "", errCredentialsNotFound
	}

	res, err := fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), credentials.Email),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
		},
	))
	if err != nil {
		log.Println("fetch [:8081/api/user/email] error")
		return "", fmt.Errorf("%w: %v", errUserNotFound, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errUserNotFound
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		return "", err
	}

	err = hash.VerifyPassword(*user.Password, credentials.Password)
	if err != nil {
		return "", err
	}

	token, err := jwt.Create(
		map[string]string{
			"email": credentials.Email,
		},
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		log.Println("Credentials Login [jwt.Create] error")
		return "", fmt.Errorf("%w: %v", errCreateToken, err)
	}
	res, err = fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/user_id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), user.Id.Hex()),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
		},
	))
	if err != nil {
		log.Println("Credentials Login fetch [:8082/api/session/user_id/{id}] error")
		return "", fmt.Errorf("%w: %v", errSessionNotRetrieved, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errSessionNotRetrieved
	}

	userSessions, err := decode.Json[[]session.SessionModel](res.Body)
	if err != nil {
		return "", err
	}

	var userSession *session.SessionModel = nil
	for _, session := range userSessions {
		if session.ProviderName == nil {
			userSession = &session
			break
		}
	}

	if userSession == nil {
		return "", errSessionNotFound
	}

	expiry, err := jwt.Expiry(token, os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", err
	}

	updateSession := session.UpdateSessionModel{
		AccessToken: &token,
		Expiry:      &expiry,
	}
	body, err := json.Marshal(updateSession)
	if err != nil {
		return "", err
	}
	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/session/id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), userSession.Id.Hex()),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		log.Println("Credentials Login fetch [:8082/api/session/id/{id}] error")
		return "", fmt.Errorf("%w: %v", errSessionNotRetrieved, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %v", errCreateuser, err)
	}
	return token, nil
}

func (m Model) Logout(ctx context.Context) (string, error) {
	// TODO: implement logout
	return "", nil
}

func (m Model) Register(regsiterModel RegisterModel) (string, error) {
	body, err := json.Marshal(regsiterModel.User)
	if err != nil {
		log.Println(err)
		return "", err
	}

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			nil,
		),
	)

	if err != nil {
		return "", fmt.Errorf("%w: %v", errCreateuser, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errCreateuser
	}

	user, err := decode.Json[user.UserModel](res.Body)
	if err != nil {
		return "", err
	}

	addSession := session.AddSessionModel{
		UserId:       user.Id,
		ProviderName: nil,
		AccessToken:  "",
		RefreshToken: nil,
		Expiry:       time.Now(),
		IdToken:      nil,
	}
	body, err = json.Marshal(addSession)
	if err != nil {
		return "", err
	}

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
		return "", fmt.Errorf("%w: %v", errCreateSession, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errCreateSession
	}

	accessToken, err := m.Login(context.WithValue(
		context.TODO(),
		CredentialsCtxKey,
		CredentialsModel{
			Email:    regsiterModel.User.Email,
			Password: *regsiterModel.User.Password,
		},
	))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (m Model) VerifyToken(token string) error {
	if err := jwt.Verify(token, os.Getenv("TOKEN_SECRET")); err == nil {
		return nil
	}

	res, err := fetch.Fetch(&http.Client{}, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/access_token/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), token),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
		}))
	if err != nil {
		return fmt.Errorf("%w: %v", errTokenNotFound, err)

	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errSessionNotRetrieved
	}
	return nil
}
