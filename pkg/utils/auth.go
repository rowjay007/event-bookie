package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtSecret should be securely stored in your environment variables
var jwtSecret = []byte("your_jwt_secret_key")

// Claims defines the custom claims for the JWT toke
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"` // Add role to claims
	jwt.StandardClaims
}


// GenerateToken generates a new JWT token for a given user ID
func GenerateToken(userID uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// Parse the token with the custom claims struct
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
