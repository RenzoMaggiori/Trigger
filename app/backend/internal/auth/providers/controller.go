package providers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
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
	_, err := h.Service.Logout(context.WithValue(context.TODO(), ProviderCtxKey, r.Header.Get("Authorization")))
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to logout", http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
