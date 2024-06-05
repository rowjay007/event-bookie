package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// GenerateReferenceID generates a unique reference ID with a prefix
func GenerateReferenceID(prefix string) (string, error) {
	uuid := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s_%s_%d", prefix, uuid, timestamp), nil
}

// Example functions to generate specific types of reference IDs
func GenerateAdminReferenceID() (string, error) {
	return GenerateReferenceID("ADMIN")
}

func GeneratePaystackReferenceID() (string, error) {
	return GenerateReferenceID("PSTK")
}
