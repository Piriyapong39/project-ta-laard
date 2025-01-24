package utils

import (
	"klui/clean-arch/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JwtSign(user models.User) (string, error) {
	if err := godotenv.Load("../config/.env"); err != nil {
		return "", err
	}
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"userId":    user.UserId,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"address":   user.Address,
		"isSeller":  user.IsSeller,
		"exp":       time.Now().Add(time.Hour * 6).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}
