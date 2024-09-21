package gmail

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Register(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusMethodNotAllowed)
	// TODO: Register user to service
}

func (h *Handler) Webhook(res http.ResponseWriter, req *http.Request) {
	var body any
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Webhook triggered, received body=%+v\n", body)
	// TODO: Handle action
	res.WriteHeader(http.StatusOK)
}
