package credentials

import (
	"net/http"
	"os"
	"time"

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

	accessToken, err := h.Service.Login(credentials)
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

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.Logout(r.Context()); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	cookie := &http.Cookie{
		Name:     authCookieName,
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   false, // TODO: true when in production
		Expires:  time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(w, cookie)
}