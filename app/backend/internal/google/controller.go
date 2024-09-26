package google

import (
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
)

type contextKey string

const providerKey contextKey = "provider"

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	// try to get the user without re-authenticating
	// provider := r.URL.Query().Get("provider")
	// r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Println(gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	// provider := r.URL.Query().Get("provider")
	// r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))
	user, err := gothic.CompleteUserAuth(w, r)

	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to complete the user auth", http.StatusUnprocessableEntity)
	}
	log.Println(user)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
}
