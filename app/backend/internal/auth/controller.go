package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"trigger.com/trigger/pkg/decode"
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
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	// Cookie or Header?
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// newUser, err := decode.Json[RegisterModel](r.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
	// 	return
	// }

	// // call user service to create user

	// accessToken, err := h.Service.Login(context.WithValue(
	// 	context.TODO(),
	// 	CredentialsCtxKey,
	// 	"",
	// ))
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
	// 	return
	// }

	// // Cookie or Header?
	// w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
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

	// // Cookie or Header?
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
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
		http.Error(w, "unable to get authorization token", http.StatusBadRequest)
		return
	}
}
