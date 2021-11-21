package util

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/masiucd/go-jwt/app/config"
)

// GenerateToken new auth token
func GenerateToken() (string, error) {
	// HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    1,
		"expire": time.Now().Add(time.Hour * 3).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
