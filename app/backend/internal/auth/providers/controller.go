package providers

import (
	"context"
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
		// redirect the user to provider oauth2 workflow
		gothic.BeginAuthHandler(w, r)
		return
	}
	accessToken, err := h.Service.Login(context.WithValue(r.Context(), LoginCtxKey, gothUser))

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	cookie := &http.Cookie{Name: authCookieName, Value: accessToken, Expires: gothUser.ExpiresAt}
	http.SetCookie(w, cookie) // TODO: fix the thing
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	// Remove the session from the database
	_, err := h.Service.Logout(context.WithValue(r.Context(), AuthorizationHeaderCtxKey, r.Header.Get("Authorization")))
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
