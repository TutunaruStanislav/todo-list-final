package api

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secretKey is the secret key you defined in the JWT_SECRET environment variable.
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// createToken - generates and returns jwt token if successful, otherwise error.
func createToken(username string) (string, error) {
	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// verifyToken - validates the jwt token and returns true if successful, otherwise false.
func verifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}

	return true
}
