package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken creates a random token for session management
func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
