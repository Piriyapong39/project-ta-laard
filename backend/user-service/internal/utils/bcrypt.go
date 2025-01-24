package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if err := godotenv.Load("../config/.env"); err != nil {
		return "", err
	}
	SALTROUND, err := strconv.Atoi(os.Getenv("SALTROUND"))
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), SALTROUND)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}
