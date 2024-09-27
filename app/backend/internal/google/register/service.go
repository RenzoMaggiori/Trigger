package register

import (
	"context"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

// TODO handle Auth, Callback, and Logout

func (m Model) Auth(ctx context.Context) (string, error) {
	return "", nil
}

func (m Model) Callback(ctx context.Context) (string, error) {

	return "", nil
}

func (m Model) Logout(ctx context.Context) (string, error) {

	return "", nil
}
