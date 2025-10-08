package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hash hashes a plaintext password using bcrypt.
func Hash(plain string) (string, error) {
	if plain == "" {
		return "", fmt.Errorf("password is empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare compares a bcrypt hash with a plaintext password.
func Compare(hashed, plain string) bool {
	if hashed == "" || plain == "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
