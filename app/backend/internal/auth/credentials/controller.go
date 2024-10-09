package credentials

import (
	"context"
	"net/http"
	"os"

	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/jwt"
)

const (
	authCookieName string = "Authorization"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	credentials, err := decode.Json[CredentialsModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	accessToken, err := h.Service.Login(context.WithValue(
		context.TODO(),
		CredentialsCtxKey,
		credentials,
	))
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	expires, err := jwt.Expiry(
		accessToken,
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	cookie := &http.Cookie{
		Name:     authCookieName,
		Value:    accessToken,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   false, // TODO: true when in production
		Expires:  expires,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	newUser, err := decode.Json[RegisterModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if newUser.User.Password == nil {
		customerror.Send(w, errCredentialsNotFound, errCodes)
		return
	}

	accessToken, err := h.Service.Register(newUser)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	expires, err := jwt.Expiry(
		accessToken,
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	cookie := &http.Cookie{
		Name:     authCookieName,
		Value:    accessToken,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   false, // TODO: true when in production
		Expires:  expires,
	}
	http.SetCookie(w, cookie)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.FromRequest(r.Header.Get("Authorization"))
	if err != nil {
		customerror.Send(w, errAuthorizationHeaderNotFound, errCodes)
		return
	}

	if err = h.Service.VerifyToken(token); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}
