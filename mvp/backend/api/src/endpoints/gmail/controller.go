package gmail

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"trigger.com/api/src/lib"
	"trigger.com/api/src/middleware"
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
	accessToken, ok := req.Context().Value(middleware.AuthHeaderCtxKey).(string)
	if !ok {
		log.Println("could not retrieve access token")
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	err := h.Gmail.Register(context.WithValue(req.Context(), gmailAccessTokenKey, accessToken))
	if err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h *Handler) Webhook(res http.ResponseWriter, req *http.Request) {
	body, err := lib.JsonDecode(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Webhook triggered, received body=%+v\n", body)
	// TODO: Handle action
	res.WriteHeader(http.StatusOK)
}
