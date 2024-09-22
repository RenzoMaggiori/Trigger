package user

import (
	"encoding/json"
	"log"
	"net/http"

	"trigger.com/api/src/lib"
)

func (h *Handler) Add(res http.ResponseWriter, req *http.Request) {
	new, err := lib.JsonDecode[User](req.Body)
	if err != nil {
		log.Println(err)
		http.Error(res, "could not parse user", http.StatusUnprocessableEntity)
		return
	}

	id, err := h.Service.Add(new)
	if err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(map[string]any{"id": id}); err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetByEmail(res http.ResponseWriter, req *http.Request) {
	email := req.PathValue("email")
	user, err := h.Service.GetByEmail(email)
	if err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(user); err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
}
