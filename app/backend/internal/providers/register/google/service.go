package google

import (
	"context"
	"fmt"

	"github.com/markbates/goth"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

// TODO handle Auth, Callback, and Logout

func (m Model) Callback(user goth.User) (string, error) {

	fmt.Println(user)
	return "", nil
}

func (m Model) Logout(ctx context.Context) (string, error) {

	return "", nil
}
