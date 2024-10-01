package credentials

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/jwt"
)

const (
	authCookieName string = "Authorization"
)

var (
	errPasswordNotFound error = errors.New("password not found")
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := decode.Json[CredentialsModel](r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	accessToken, err := h.Service.Login(context.WithValue(
		context.TODO(),
		CredentialsCtxKey,
		credentials,
	))
	if err != nil {
		log.Println(err)
		http.Error(w, "eror while retrieving token", http.StatusNotFound)
		return
	}

	expires, err := jwt.Expiry(
		accessToken,
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Name: authCookieName, Value: accessToken, Expires: expires}
	http.SetCookie(w, cookie)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	newUser, err := decode.Json[RegisterModel](r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}
	if newUser.User.Password == nil {
		log.Println(errPasswordNotFound)
		http.Error(w, errPasswordNotFound.Error(), http.StatusUnprocessableEntity)
		return
	}

	accessToken, err := h.Service.Register(newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to login user", http.StatusInternalServerError)
		return
	}

	expires, err := jwt.Expiry(
		accessToken,
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Name: authCookieName, Value: accessToken, Expires: expires}
	http.SetCookie(w, cookie)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	token, err := h.Service.GetToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to get authorization token", http.StatusBadRequest)
		return
	}

	if err = h.Service.VerifyToken(token); err != nil {
		log.Println(err)
		http.Error(w, "unable to retrieve the token from the db", http.StatusBadRequest)
		return
	}
}
