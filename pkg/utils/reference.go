package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateReferenceID(prefix string) (string, error) {
	uuid := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s_%s_%d", prefix, uuid, timestamp), nil
}

func GenerateAdminReferenceID() (string, error) {
	return GenerateReferenceID("ADMIN")
}

func GeneratePaystackReferenceID() (string, error) {
	return GenerateReferenceID("PSTK")
}
