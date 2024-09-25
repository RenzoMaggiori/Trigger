package mongodb

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 14

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
