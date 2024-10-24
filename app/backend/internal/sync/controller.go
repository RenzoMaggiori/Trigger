package sync

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth/gothic"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SyncWith(w http.ResponseWriter, r *http.Request) {
	access_token := r.Header.Get("Authorization")

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		gothic.BeginAuthHandler(w, r)
		return
	}

	err = h.Service.SyncWith(gothUser, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s/settings", os.Getenv("WEB_BASE_URL")), http.StatusPermanentRedirect)
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	access_token := r.Header.Get("Authorization")

	user, ko := gothic.CompleteUserAuth(w, r)
	if ko != nil {
		http.Error(w, "failed to complete user auth", http.StatusInternalServerError)
		return
	}

	err := h.Service.Callback(user, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s/settings", os.Getenv("WEB_BASE_URL")), http.StatusPermanentRedirect)
}
