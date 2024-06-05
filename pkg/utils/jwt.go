package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"` 
	jwt.StandardClaims
}


func GenerateJWT(email, role string) (string, error) {
	claims := &Claims{
		Email: email,
		Role:  role, 
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func GenerateResetToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := bcrypt.GenerateFromPassword(tokenBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(tokenBytes), nil
}

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
