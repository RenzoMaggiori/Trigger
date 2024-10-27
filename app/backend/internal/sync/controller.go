package sync

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/markbates/goth/gothic"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SyncWith(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		h.Service.GrantAccess(w, r)
		return
	}

	access_token := r.Header.Get("Authorization")
	err = h.Service.SyncWith(gothUser, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	split := strings.Split(state, ":")
	url := split[0]
	token := split[1]

	urlDecodedBytes, err := base64.URLEncoding.DecodeString(url)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	redirectUrl := string(urlDecodedBytes)
	log.Println(redirectUrl)

	tokenDecodedBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	access_token := string(tokenDecodedBytes)
	log.Println(access_token)

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
