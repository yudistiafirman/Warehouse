package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Fungsi buat generate token JWT
func GenerateJWT(userID uint, level string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"level":   level,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expired dalam 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
