package hash

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
