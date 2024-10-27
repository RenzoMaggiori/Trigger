package sync

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SyncWith(w http.ResponseWriter, r *http.Request) {
	access_token := r.Header.Get("Authorization")
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		h.Service.GrantAccess(w, r)
		return
	}

	err = h.Service.SyncWith(gothUser, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	access_token := r.Header.Get("Authorization")

	state := r.URL.Query().Get("state")
	stateDecodedBytes, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	redirectUrl := string(stateDecodedBytes)
	log.Println(redirectUrl)
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.Callback(user, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	http.Redirect(w, r, redirectUrl, http.StatusPermanentRedirect)
}
