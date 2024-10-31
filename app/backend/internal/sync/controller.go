package sync

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SyncWith(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		h.Service.GrantAccess(w, r)
		return
	}

	access_token := r.URL.Query().Get("token")
	err = h.Service.SyncWith(gothUser, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	stateArr := strings.Split(state, ":")
	if len(stateArr) != 2 {
		customerror.Send(w, errors.ErrMalformedState, errors.ErrCodes)
		return
	}

	url := stateArr[0]
	token := stateArr[1]
	urlDecodedBytes, err := base64.URLEncoding.DecodeString(url)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	redirectUrl := string(urlDecodedBytes)
	tokenDecodedBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	access_token := string(tokenDecodedBytes)
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	// TODO: Implement this FOR DISCORD ONLY
	// guildID := r.URL.Query().Get("guild_id")

	err = h.Service.Callback(user, access_token)
	if err != nil {
		log.Println(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	http.Redirect(w, r, redirectUrl, http.StatusPermanentRedirect)
}

func (h *Handler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))
	provider := r.PathValue("provider")

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadUserId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	sync, err := h.Service.ByUserId(userId, provider)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, sync); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}
