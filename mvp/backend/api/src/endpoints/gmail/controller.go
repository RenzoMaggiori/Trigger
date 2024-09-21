package gmail

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) AuthProvider(res http.ResponseWriter, req *http.Request) {
	authUrl := h.Gmail.Provider()
	http.Redirect(res, req, authUrl, http.StatusPermanentRedirect)
}

func (h *Handler) AuthCallback(res http.ResponseWriter, req *http.Request) {
	// TODO: get authcode from request
	token, err := h.Gmail.Callback(req.Context(), "")

	if err != nil {
		http.Error(res, "Unable to authenticate", http.StatusUnauthorized)
		return
	}
	res.Header().Set("Authentication", fmt.Sprintf("Bearer %s", token.AccessToken))
	res.WriteHeader(http.StatusOK)
}

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
