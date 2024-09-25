package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := decode.Json[LoginModel](r.Body)
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
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	// Cookie or Header?
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	newUser, err := decode.Json[RegisterModel](r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	// call user service to create user

	accessToken, err := h.Service.Login(context.WithValue(
		context.TODO(),
		CredentialsCtxKey,
		"",
	))
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	// Cookie or Header?
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}
