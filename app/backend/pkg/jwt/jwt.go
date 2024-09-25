package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errTokenNotValid error = errors.New("token is not valid")
)

func Create(claims map[string]string, secret string) (string, error) {
	tokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	for k, v := range claims {
		tokenClaims[k] = v
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		tokenClaims,
	)
	rawToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return rawToken, nil
}

func Verify(rawToken string, secret string) error {
	token, err := jwt.Parse(rawToken, func(_ *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errTokenNotValid
	}
	return nil
}
