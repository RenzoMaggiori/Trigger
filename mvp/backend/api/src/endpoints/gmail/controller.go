package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (h *Handler) AuthProvider(res http.ResponseWriter, req *http.Request) {
	authUrl := h.Gmail.Provider(res)
	http.Redirect(res, req, authUrl, http.StatusTemporaryRedirect)
}

func (h *Handler) AuthCallback(res http.ResponseWriter, req *http.Request) {
	// TODO: store refresh token in database
	token, err := h.Gmail.Callback(req)
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}

	authCookie := http.Cookie{Name: "access_token", Value: token.AccessToken, Expires: time.Unix(token.ExpiresIn, 0)}
	http.SetCookie(res, &authCookie)
	http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
}

func (h *Handler) Register(res http.ResponseWriter, req *http.Request) {
	accessHeader := strings.Split(req.Header.Get("Authorization"), " ")
	if len(accessHeader) < 2 {
		http.Error(res, "could not retrieve access token from header", http.StatusBadRequest)
		return
	}

	err := h.Gmail.Register(context.WithValue(req.Context(), gmailAccessTokenKey, accessHeader[1]))
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
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
