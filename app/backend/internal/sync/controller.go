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
		log.Println("redirecting...")
		gothic.BeginAuthHandler(w, r)
		return
	}

	log.Println("syn with...")
	err = h.Service.SyncWith(gothUser)

	if err != nil {
		// customerror.Send(w, err, errCodes)
		http.Error(w, "failed to sync with user", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	user, ko := gothic.CompleteUserAuth(w, r)
	if ko != nil {
		//! ret error
		// customerror.Send(w, err, errCodes)
		http.Error(w, "failed to complete user auth", http.StatusInternalServerError)
		return
	}

	//* : we might want to get the cookie here and pass it to the service so we can get de user logged in
	// cookie, err := r.Cookie("Authorization")
	// if err != nil {
	// 	log.Println("failed to get cookie")
		// http.Error(w, "failed to get cookie", http.StatusInternalServerError)
		// return
	// }

	// Store the user and the session in the database
	err := h.Service.Callback(user)
	if err != nil {
		//! ret error
		// customerror.Send(w, err, errCodes)
		http.Error(w, "failed to store sync in the database", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("%s/settings", os.Getenv("WEB_BASE_URL")), http.StatusPermanentRedirect)
}
