package providers

import (
	"encoding/base64"
	"net/http"
	"os"
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/user"
)

func (m Model) Login(w http.ResponseWriter, r *http.Request) error {
	redirectUrl := r.URL.Query().Get("redirect")
	log.Println(redirectUrl)
	state := base64.URLEncoding.EncodeToString([]byte(redirectUrl))
	values := r.URL.Query()
	values.Set("state", state)
	r.URL.RawQuery = values.Encode()
	gothic.BeginAuthHandler(w, r)
	return nil
}

func (m Model) AccessToken(gothUser goth.User) (string, error) {
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
		accesToken, err := m.AccessToken(gothUser)
		if err != nil {
			return "", err
		}
		return accesToken, nil
	}
	return "", errUserNotFound
}

func (m Model) Logout(w http.ResponseWriter, r *http.Request) error {
	// TODO: implement logout
	gothic.Logout(w, r)
	return nil
}
