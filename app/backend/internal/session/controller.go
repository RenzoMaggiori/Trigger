package session

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetSessions(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetSessionById(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetSessionByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetByUserId(user_id)
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

func (h *Handler) AddSession(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddSessionModel](r.Body)
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

func (h *Handler) UpdateSessionById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}

	update, err := decode.Json[UpdateSessionModel](r.Body)
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

func (h *Handler) UpdateSessionByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))
	providerName := r.PathValue("providerName")

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}
	update, err := decode.Json[UpdateSessionModel](r.Body)

	if err != nil {
		log.Print(err)
		http.Error(w, "could not decode json", http.StatusUnprocessableEntity)
		return
	}

	updatedUser, err := h.Service.UpdateByUserId(userId, providerName, &update)
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

func (h *Handler) DeleteSessionById(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) DeleteSessionByUserId(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))
	providerName := r.PathValue("providerName")

	if err != nil {
		log.Print(err)
		http.Error(w, "bad user id", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteByUserId(userId, providerName); err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
