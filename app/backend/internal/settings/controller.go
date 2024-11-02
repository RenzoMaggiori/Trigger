package settings

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/session"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	sync, err := h.Service.GetById(id)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, sync); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}

func (h *Handler) GetMySettings(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}

	s, _, err := session.GetSessionByAccessTokenRequest(token)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	sync, err := h.Service.GetByUserId(s.UserId)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, sync); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}

func (h *Handler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	sync, err := h.Service.GetByUserId(id)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, sync); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var add AddSettingsModel
	if err := json.NewDecoder(r.Body).Decode(&add); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err := h.Service.Add(&add); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var update UpdateSettingsModel
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err := h.Service.Update(&update); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}
