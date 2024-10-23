package sync

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth/gothic"
)

func (h *Handler) SyncWith(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		// redirect the user to provider oauth2 workflow
		gothic.BeginAuthHandler(w, r)
		return
	}

	err = h.Service.SyncWith(gothUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	user, ko := gothic.CompleteUserAuth(w, r)
	if ko != nil {
		http.Error(w, "failed to complete user auth", http.StatusInternalServerError)
		return
	}
	err := h.Service.Callback(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("%s/settings", os.Getenv("WEB_BASE_URL")), http.StatusPermanentRedirect)
}
