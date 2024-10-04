package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) Login(ctx context.Context) (string, error) {
	gothUser, ok := ctx.Value(LoginCtxKey).(goth.User)

	if !ok {
		return "", errCredentialsNotFound
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/user/email/%s", os.Getenv("USER_SERVICE_BASE_URL"), gothUser.Email),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return "", errUserNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errUserNotFound
	}
	user, err := decode.Json[user.UserModel](res.Body)

	if err != nil {
		return "", errUserTypeNone
	}

	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodGet,
			fmt.Sprintf("%s/api/session/user_id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), user.Id.Hex()),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return "", errSessionNotFound
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errSessionNotFound
	}
	sessions, err := decode.Json[[]session.SessionModel](res.Body)

	if err != nil {
		return "", errSessionTypeNone
	}
	var providerSession *session.SessionModel = nil
	for _, s := range sessions {
		if *s.ProviderName == gothUser.Provider {
			providerSession = &s
		}
	}
	if providerSession == nil {
		return "", errProviderSessionNotFound
	}

	patchSession := session.UpdateSessionModel{
		AccessToken:  &gothUser.AccessToken,
		RefreshToken: &gothUser.RefreshToken,
		Expiry:       &gothUser.ExpiresAt,
		IdToken:      &gothUser.IDToken,
	}

	body, err := json.Marshal(patchSession)

	if err != nil {
		return "", err
	}

	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/api/session/id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), providerSession.Id.Hex()),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return "", errSessionPatchFailed
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errSessionPatchFailed
	}

	return gothUser.AccessToken, nil
}

func (m Model) Callback(gothUser goth.User) (string, error) {
	addUser := user.AddUserModel{
		Email:    gothUser.Email,
		Password: nil,
	}

	body, err := json.Marshal(addUser)
	if err != nil {
		return "", err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/user/add", os.Getenv("USER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		return "", err
	}
	if res.StatusCode == http.StatusOK {
		user, err := decode.Json[user.UserModel](res.Body)
		if err != nil {
			return "", err
		}

		addSession := session.AddSessionModel{
			UserId:       user.Id,
			ProviderName: &gothUser.Provider,
			AccessToken:  gothUser.AccessToken,
			RefreshToken: &gothUser.RefreshToken,
			Expiry:       gothUser.ExpiresAt,
			IdToken:      &gothUser.IDToken,
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
				nil,
			),
		)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return "", errors.New("unable to create session")
		}
		return gothUser.AccessToken, nil
	}

	if res.StatusCode == http.StatusConflict {
		accesToken, err := m.Login(context.WithValue(context.TODO(), LoginCtxKey, gothUser))

		if err != nil {
			return "", err
		}
		return accesToken, nil
	}
	return "", errUserNotFound
}

func (m Model) Logout(ctx context.Context) (string, error) {
	accessToken, ok := ctx.Value(AuthorizationHeaderCtxKey).(string)

	_ = accessToken
	if !ok {
		return "", errCredentialsNotFound
	}
	// TODO: implement logout
	return "", nil
}
