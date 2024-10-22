package authenticator

import "context"

type AuthorizationCtx string

const AuthorizationTokenCtxKey AuthorizationCtx = AuthorizationCtx("AuthorizationCtxKey")

type Authenticator interface {
	Login(ctx context.Context) (string, error)
	Logout(ctx context.Context) (string, error)
}
