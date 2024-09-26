package google

import (
	"context"
	"fmt"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func (m Model) Login(ctx context.Context) (string, error) {

	fmt.Println("ASdasdsad")
	return "", nil
}
