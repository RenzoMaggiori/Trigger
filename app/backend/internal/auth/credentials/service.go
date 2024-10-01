package credentials

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/hash"
	"trigger.com/trigger/pkg/jwt"
)

var (
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find token in authorization header")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
	errUserNotRetrieved            error = errors.New("could not retrieve user")
	errSessionNotRetrieved         error = errors.New("could not retrieve session")
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
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errUserNotRetrieved
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
	res, err = fetch.Fetch(http.DefaultClient, fetch.NewFetchRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/session/user_id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), user.Id.String()),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
		},
	))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", errSessionNotRetrieved
	}

	userSession, err := decode.Json[session.SessionModel](res.Body)
	if err != nil {
		return "", err
	}

	expiry, err := jwt.Expiry(token, os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", err
	}

	updateSession := session.UpdateSessionModel{
		AccessToken: &token,
		Expiry:      &expiry,
	}
	body, err := bson.Marshal(updateSession)
	if err != nil {
		return "", err
	}

	res, err = fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodPatch,
			fmt.Sprintf("%s/session/user_id/%d", os.Getenv("USER_SERVICE_BASE_URL"), userSession.Id),
			bytes.NewReader(body),
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("ADMIN_TOKEN")),
			},
		),
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", res.Status)
		return "", errors.New("unable to create user")
	}
	return token, nil
}

func (m Model) Logout(ctx context.Context) (string, error) {
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
		log.Println(err)
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", res.Status)
		return "", errors.New("unable to create user")
	}

	user, err := decode.Json[user.UserModel](res.Body)

	if err != nil {
		log.Println(err)
		return "", err
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
	expiry, err := jwt.Expiry(accessToken)
	if err != nil {
		return "", err
	}
	addSession := session.AddSessionModel{
		UserId:       user.Id,
		ProviderName: nil,
		AccessToken:  accessToken,
		RefreshToken: nil,
		Expiry:       expiry,
		IdToken:      nil,
	}

	body, err = json.Marshal(addSession)
	if err != nil {
		log.Println(err)
		return "", err
	}

	res, err = fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/session/add", os.Getenv("SESSION_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			nil,
		),
	)

	if err != nil {
		log.Println(err)
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", res.Status)
		return "", errors.New("unable to create session")
	}

	return accessToken, nil
}

func (m Model) GetToken(authorizationHeader string) (string, error) {
	if authorizationHeader == "" {
		return "", errAuthorizationHeaderNotFound
	}

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			return "", errTokenNotFound
		}
		m.authType = Credentials
		return parts[1], nil
	}
	// TODO: check for oauth token
	return "", errAuthorizationTypeNone
}

func (m Model) VerifyToken(token string) error {
	switch m.authType {
	case Credentials:
		return jwt.Verify(token, os.Getenv("TOKEN_SECRET"))
	case OAuth:
		// TODO: verify oauth2 token
		return errAuthTypeUndefined
	default:
		return errAuthTypeUndefined
	}
}
