package worker

import (
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/jwt"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.FromRequest(r.Header.Get("Authorization"))
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	me, err := h.Service.Me(token)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}

	if err = encode.Json(w, me); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}


func (h *Handler) GetGuildChannels(w http.ResponseWriter, r *http.Request) {
	guildID := r.PathValue("guild_id")

	if guildID == "" {
		customerror.Send(w, errors.ErrUserTypeNone, errors.ErrCodes)
		return
	}

	ch, err := h.Service.GuildChannels(guildID)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}

	if err = encode.Json(w, ch); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}


}