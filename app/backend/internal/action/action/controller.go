package action

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetActions(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.Get()

	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, users); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetActionById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadActionId, err)
		customerror.Send(w, error, errCodes)
		return
	}

	user, err := h.Service.GetById(id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetActionsByProvider(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")

	user, err := h.Service.GetByProvider(provider)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}
func (h *Handler) GetActionByAction(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")

	user, err := h.Service.GetByAction(action)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, user); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) AddAction(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddActionModel](r.Body)

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
