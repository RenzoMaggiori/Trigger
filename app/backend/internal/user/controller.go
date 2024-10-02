package user

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.Get()

	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, users); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetById(id)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, user); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	user, err := h.Service.GetByEmail(email)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, user); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddUserModel](r.Body)

	if err != nil {
		log.Print(err)
		http.Error(w, "could not decode json", http.StatusUnprocessableEntity)
		return
	}

	newUser, err := h.Service.Add(&add)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, newUser); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}

	update, err := decode.Json[UpdateUserModel](r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "could not decode json", http.StatusUnprocessableEntity)
		return
	}

	updatedUser, err := h.Service.UpdateById(id, &update)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, updatedUser); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	update, err := decode.Json[UpdateUserModel](r.Body)

	if err != nil {
		log.Print(err)
		http.Error(w, "could not decode json", http.StatusUnprocessableEntity)
		return
	}

	updatedUser, err := h.Service.UpdateByEmail(email, &update)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err = encode.Json(w, updatedUser); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteById(id); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")

	if err := h.Service.DeleteByEmail(email); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
