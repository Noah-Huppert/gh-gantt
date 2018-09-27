package encryption

import (
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// ComputeKey creates a 32 byte encryption key from a password
func ComputeKey(password string) ([]byte, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error running bcrypt on password: %s", err.Error())
	}

	hashed := sha512.Sum512_256(key)
	return hashed[:], nil
}
