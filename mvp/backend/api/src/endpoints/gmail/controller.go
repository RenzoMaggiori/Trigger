package gmail

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"trigger.com/api/src/lib"
	"trigger.com/api/src/middleware"
)

func (h *Handler) AuthProvider(res http.ResponseWriter, req *http.Request) {
	authUrl := h.Service.Provider(res)
	http.Redirect(res, req, authUrl, http.StatusTemporaryRedirect)
}

func (h *Handler) AuthCallback(res http.ResponseWriter, req *http.Request) {
	// Get token
	token, err := h.Service.Callback(req)
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}

	// Get user from google
	gmailUser, err := h.Service.GetUser(token)
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}

	// Check if user exists
	getUserRes, err := lib.Fetch(lib.NewFetchRequest(
		"GET",
		fmt.Sprintf("%s/user/%s", os.Getenv("API_URL"), gmailUser.EmailAddress),
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
		},
	))
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}
	defer getUserRes.Body.Close()
	if getUserRes.StatusCode == http.StatusOK {
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
	}

	// Add user to db
	addUserRes, err := lib.Fetch(lib.NewFetchRequest(
		"POST",
		fmt.Sprintf("%s/user", os.Getenv("API_URL")),
		map[string]any{
			// TODO: add email
			"email":        gmailUser.EmailAddress,
			"accessToken":  token.AccessToken,
			"refreshToken": token.RefreshToken,
			"tokenType":    token.TokenType,
			"expiry":       token.Expiry,
		},
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
		},
	))
	if err != nil {
		log.Println(err)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return
	}
	defer addUserRes.Body.Close()
	if addUserRes.StatusCode != http.StatusOK {
		log.Printf("invalid status code, received %s\n", addUserRes.Status)
		http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
		return

	}

	http.Redirect(res, req, fmt.Sprintf("%s/", os.Getenv("WEB_URL")), http.StatusPermanentRedirect)
}

func (h *Handler) Register(res http.ResponseWriter, req *http.Request) {
	accessToken, ok := req.Context().Value(middleware.AuthHeaderCtxKey).(string)
	if !ok {
		log.Println("could not retrieve access token")
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	err := h.Service.Register(context.WithValue(req.Context(), gmailAccessTokenKey, accessToken))
	if err != nil {
		log.Println(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h *Handler) Webhook(res http.ResponseWriter, req *http.Request) {
	body, err := lib.JsonDecode[Event](req.Body)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Webhook triggered, received body=%+v\n", body)
	h.Service.Webhook(context.WithValue(req.Context(), gmailEventKey, body))
	res.WriteHeader(http.StatusOK)
}
