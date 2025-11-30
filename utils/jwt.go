package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define a custom claims structure (optional but recommended)
type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(userID int, secretKey string) (string, error) {
	// Set the claims including standard registered claims
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Set expiration time (e.g., 15 minutes)
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app-issuer",
		},
	}

	// Create the token with the specified signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HMAC SHA256

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
func VerifyToken(tokenString, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		// Provide the secret key for validation
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing/validating token: %w", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token claims")
	}
}
