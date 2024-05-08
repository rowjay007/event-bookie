package utils

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key")

// Claims represents the claims contained within the JWT
type Claims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

// GenerateToken generates a new JWT token
func GenerateToken(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// VerifyToken verifies the JWT token and returns the claims if valid
func VerifyToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, jwt.ErrSignatureInvalid
}
