package session

import (
	"net/http"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetSessions(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.Get()

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, users); err != nil {
		customerror.Send(w, err, errCodes)																						
		return
	}
}

func (h *Handler) GetSessionById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadSessionID, err)
		customerror.Send(w, error, errCodes)
		return
	}

	user, err := h.Service.GetById(id)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetSessionByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadUserID, err)
		customerror.Send(w, error, errCodes)
		return
	}

	user, err := h.Service.GetByUserId(userId)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) AddSession(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddSessionModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	newUser, err := h.Service.Add(&add)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, newUser); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) UpdateSessionById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadSessionID, err)
		customerror.Send(w, error, errCodes)
		return
	}

	update, err := decode.Json[UpdateSessionModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	updatedUser, err := h.Service.UpdateById(id, &update)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, updatedUser); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) DeleteSessionById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadSessionID, err)
		customerror.Send(w, error, errCodes)
		return
	}
	if err := h.Service.DeleteById(id); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetByToken(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("access_token")

	session, err := h.Service.GetByToken(token)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, session); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}
