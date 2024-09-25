package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/fetch"
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
	newUser, err := decode.Json[RegisterModel](r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	body, err := json.Marshal(newUser.User)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess user", http.StatusUnprocessableEntity)
		return
	}

	res, err := fetch.Fetch(
		&http.Client{},
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("%s/api/user", os.Getenv("USER_SERVICE_BASE_URL")),
			bytes.NewReader(body),
			nil,
		),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to create user", http.StatusInternalServerError)
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", res.Status)
		http.Error(w, "unable to create user", http.StatusBadRequest)
		return
	}

	accessToken, err := h.Service.Login(context.WithValue(
		context.TODO(),
		CredentialsCtxKey,
		CredentialsModel{
			Email:    newUser.User.Email,
			Password: *newUser.User.Password,
		},
	))
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to proccess body", http.StatusUnprocessableEntity)
		return
	}

	// Cookie or Header?
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}
