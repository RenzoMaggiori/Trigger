package providers

import (
	"context"
	"net/http"
	"os"

	"github.com/markbates/goth"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
)

func (m Model) Login(ctx context.Context) (string, error) {
	gothUser, ok := ctx.Value(LoginCtxKey).(goth.User)

	if !ok {
		return "", errCredentialsNotFound
	}
	user, _, err := user.GetUserByEmailRequest(os.Getenv("ADMIN_TOKEN"), gothUser.Email)

	if err != nil {
		return "", err
	}

	userSessions, _, err := session.GetSessionByUserIdRequest(os.Getenv("ADMIN_TOKEN"), user.Id.Hex())

	if err != nil {
		return "", err
	}

	var providerSession *session.SessionModel = nil
	for _, s := range userSessions {
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

	updatedSession, _, err := session.UpdateSessionByIdRequest(os.Getenv("ADMIN_TOKEN"), providerSession.Id.Hex(), patchSession)

	if err != nil {
		return "", err
	}

	return updatedSession.AccessToken, nil
}

func (m Model) Callback(gothUser goth.User) (string, error) {
	addUser := user.AddUserModel{
		Email:    gothUser.Email,
		Password: nil,
	}

	user, code, err := user.AddUserRequest(os.Getenv("ADMIN_TOKEN"), addUser)
	if code == http.StatusOK {
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
		_, _, err := session.AddSessionRequest(os.Getenv("ADMIN_TOKEN"), addSession)
		if err != nil {
			return "", err
		}
		return gothUser.AccessToken, nil
	}

	if code == http.StatusConflict {
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
