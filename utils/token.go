package utils

import (
	"os"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
