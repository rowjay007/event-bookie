package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateReferenceID(prefix string) (string, error) {
	code, err := generateRandomString(8)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s", prefix, code), nil
}

func GenerateAdminReferenceID() (string, error) {
	return GenerateReferenceID("ADMIN")
}

func GeneratePaystackReferenceID() (string, error) {
	return GenerateReferenceID("PST")
}

func GenerateFlutterwaveReferenceID() (string, error) {
	return GenerateReferenceID("FLW")
}

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		randomByte, err := randByte(charset)
		if err != nil {
			return "", err
		}
		b[i] = randomByte
	}
	return string(b), nil
}

func randByte(charset string) (byte, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}
	return charset[idx.Int64()], nil
}
