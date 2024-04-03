package helpers

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Hashed password error - %v", err)
	}
	return string(hashedPassword), nil
}



func GenerateId() uuid.UUID {
	return uuid.New()
}
