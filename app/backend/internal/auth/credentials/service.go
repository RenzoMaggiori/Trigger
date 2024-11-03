package credentials

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
	"trigger.com/trigger/pkg/fetch"
	"trigger.com/trigger/pkg/hash"
	"trigger.com/trigger/pkg/jwt"
	"trigger.com/trigger/pkg/middleware"
)

func (m Model) Login(credentials CredentialsModel) (string, error) {
	user, _, err := user.GetUserByEmailRequest(os.Getenv("ADMIN_TOKEN"), credentials.Email)
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
		return "", fmt.Errorf("%w: %v", errCreateToken, err)
	}
	userSessions, _, err := session.GetSessionByUserIdRequest(token, user.Id.Hex())

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

	session, _, err := session.UpdateSessionByIdRequest(token, userSession.Id.Hex(), updateSession)
	if err != nil {
		return "", err
	}

	return session.AccessToken, nil
}

func (m Model) Register(regsiterModel RegisterModel) (string, error) {
	addUser := user.AddUserModel{
		Email:    regsiterModel.User.Email,
		Password: regsiterModel.User.Password,
	}
	user, _, err := user.AddUserRequest(os.Getenv("ADMIN_TOKEN"), addUser)

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

	_, _, err = session.AddSessionRequest(os.Getenv("ADMIN_TOKEN"), addSession)

	if err != nil {
		return "", err
	}

	accessToken, err := m.Login(CredentialsModel{
		Email:    regsiterModel.User.Email,
		Password: *regsiterModel.User.Password,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (m Model) VerifyToken(token string) error {
	if err := jwt.Verify(token, os.Getenv("TOKEN_SECRET")); err == nil {
		return nil
	}

	_, _, err := session.GetSessionByAccessTokenRequest(token)

	if err != nil {
		return err
	}

	return nil
}

func (m Model) Logout(ctx context.Context) error {
	accessToken, ok := ctx.Value(middleware.TokenCtxKey).(string)
	if !ok {
		return errTokenNotFound
	}
	s, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		http.DefaultClient,
		fetch.NewFetchRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/api/session/id/%s", os.Getenv("SESSION_SERVICE_BASE_URL"), s.Id.Hex()),
			nil,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", accessToken),
			},
		),
	)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
