package providers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	url, err := gothic.GetAuthURL(w, r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to complete the user auth", http.StatusUnprocessableEntity)
	}
	// Store the user and the session in the database
	_, err = h.Service.Callback(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to sotre user auth", http.StatusUnprocessableEntity)
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	// Remove the session from the database
	h.Service.Logout(r.Context())
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
