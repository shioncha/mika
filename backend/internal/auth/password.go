package auth

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func NormarlizeEmail(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	return email
}

func GenerateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}
