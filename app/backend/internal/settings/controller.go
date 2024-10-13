package settings

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "trigger.com/trigger/pkg/encode"
)

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	sync, err := h.Service.GetById(id)

	if err != nil {
		http.Error(w, "could not get sync", http.StatusInternalServerError)
		return
	}

	// if err = encode.Json(w, sync); err != nil {
	// 	http.Error(w, "could not encode sync", http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sync); err != nil {
        http.Error(w, "could not encode sync", http.StatusInternalServerError)
        return
    }
	
}

func (h *Handler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	sync, err := h.Service.GetByUserId(id)

	if err != nil {
		http.Error(w, "could not get sync", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sync); err != nil {
		http.Error(w, "could not encode sync", http.StatusInternalServerError)
		return
	}
	
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var add AddSettingsModel
	if err := json.NewDecoder(r.Body).Decode(&add); err != nil {
		http.Error(w, "could not decode add settings", http.StatusBadRequest)
		return
	}

	if err := h.Service.Add(&add); err != nil {
		http.Error(w, "could not add settings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}