package worker

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	accessToken, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}

	me, err := h.Service.Me(accessToken)
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
		customerror.Send(w, errors.ErrGuildIdNotFound, errors.ErrCodes)
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

func (h *Handler) AddDiscordSession(w http.ResponseWriter, r *http.Request) {
	session, err := decode.Json[AddDiscordSessionModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.AddSession(&session)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) UpdateDiscordSession(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		customerror.Send(w, errors.ErrBadUserId, errors.ErrCodes)
		return
	}

	session, err := decode.Json[UpdateDiscordSessionModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.UpdateSession(id.Hex(), &session)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) GetCurrentDiscordSession(w http.ResponseWriter, r *http.Request) {
	accessToken, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}

	session, err := h.Service.GetSession(accessToken)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, session); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) DeleteDiscordSession(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		customerror.Send(w, errors.ErrBadUserId, errors.ErrCodes)
		return
	}

	err = h.Service.DeleteSession(id.Hex())
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}