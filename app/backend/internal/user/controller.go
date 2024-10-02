package user

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadUserID, err)
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

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	user, err := h.Service.GetByEmail(email)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddUserModel](r.Body)

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

func (h *Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadUserID, err)
		customerror.Send(w, error, errCodes)
		return
	}

	update, err := decode.Json[UpdateUserModel](r.Body)
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

func (h *Handler) UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	update, err := decode.Json[UpdateUserModel](r.Body)

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	updatedUser, err := h.Service.UpdateByEmail(email, &update)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, updatedUser); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadUserID, err)
		customerror.Send(w, error, errCodes)
		return
	}
	if err := h.Service.DeleteById(id); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) DeleteUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	if err := h.Service.DeleteByEmail(email); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}
