package providers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
	customerror "trigger.com/trigger/pkg/custom-error"
)

const (
	authCookieName string = "Authorization"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		h.Service.Login(w, r)
		return
	}

	accessToken, err := h.Service.AccessToken(gothUser)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	cookie := &http.Cookie{
		Name:     authCookieName,
		Value:    accessToken,
		Expires:  gothUser.ExpiresAt,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   false, // TODO: true when in production
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	stateDecodedBytes, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	redirectUrl := string(stateDecodedBytes)
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	// Store the user and the session in the database
	accessToken, err := h.Service.Callback(user)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	redirectWithToken := fmt.Sprintf("%s?token=%s", redirectUrl, accessToken)
	http.Redirect(w, r, redirectWithToken, http.StatusPermanentRedirect)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// Remove the session from the database
	err := h.Service.Logout(w, r)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
