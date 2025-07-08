package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateSessionToken generates a unique and secure session token
func GenerateSessionToken(byteLength int) (string, error) {
	bytes := make([]byte, byteLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes for token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
