package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length uint32) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}
