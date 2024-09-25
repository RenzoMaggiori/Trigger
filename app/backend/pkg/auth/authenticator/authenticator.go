package authenticator

import "context"

type Authenticator interface {
	Login(ctx context.Context) (string, error)
}
