package authenticator

import (
	"net/http"
)

type AuthorizationCtx string

const AuthorizationTokenCtxKey AuthorizationCtx = AuthorizationCtx("AuthorizationCtxKey")

type Authenticator interface {
	Login(w http.ResponseWriter, r *http.Request) error
	Logout(w http.ResponseWriter, r *http.Request) error
}
