package worker

import (
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/jwt"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.FromRequest(r.Header.Get("Authorization"))
	err := h.Service.Me()

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) GetGuilds(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Guilds()

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) GetGuildChannels(w http.ResponseWriter, r *http.Request) {
	guildID := r.PathValue("guild_id")

	if guildID == "" {
		customerror.Send(w, errors.ErrUserTypeNone, errors.ErrCodes)
		return
	}

	err := h.Service.GuildChannels(guildID)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}