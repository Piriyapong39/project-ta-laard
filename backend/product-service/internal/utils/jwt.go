package utils

import (
	"fmt"
	"os"
	"product-service/internal/models"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func DecodedJWT(tokenString string) (user models.User, err error) {

	if err := godotenv.Load("../config/.env"); err != nil {
		return user, err
	}

	tokenPart := strings.Split(tokenString, " ")
	if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
		return user, fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(tokenPart[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return user, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, fmt.Errorf("invalid claims")
	}

	id, ok1 := claims["userId"].(float64)
	email, ok2 := claims["email"].(string)
	firstName, ok3 := claims["firstName"].(string)
	lastName, ok4 := claims["lastName"].(string)
	isSeller, ok5 := claims["isSeller"].(bool)

	if !(ok1 && ok2 && ok3 && ok4 && ok5) {
		return user, fmt.Errorf("invalid claim values")
	}
	return models.User{
		UserId:    uint(id),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		IsSeller:  isSeller,
	}, nil
}
