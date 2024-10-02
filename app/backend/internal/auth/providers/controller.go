package providers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
)

const (
	authCookieName string = "Authorization"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		// redirect the user to provider oauth2 workflow
		log.Println(err)
		gothic.BeginAuthHandler(w, r)
		return
	}
	accessToken, err := h.Service.Login(context.WithValue(r.Context(), LoginCtxKey, gothUser))

	if err != nil {
		log.Println(err)
		http.Error(w, "unable to login", http.StatusUnprocessableEntity)
		return
	}

	cookie := &http.Cookie{Name: authCookieName, Value: accessToken, Expires: gothUser.ExpiresAt}
	http.SetCookie(w, cookie) // TODO: fix the thing
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to complete the user auth", http.StatusUnprocessableEntity)
		return
	}

	// Store the user and the session in the database
	accessToken, err := h.Service.Callback(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to store user auth", http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	// Remove the session from the database
	_, err := h.Service.Logout(context.WithValue(r.Context(), AuthorizationHeaderCtxKey, r.Header.Get("Authorization")))
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to logout", http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
