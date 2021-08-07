package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordSalt() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func GeneratePassword() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func HashPassword(salt, password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(hashedPassword)
}
