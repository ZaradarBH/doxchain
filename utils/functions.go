package utils

import (
	"crypto/rand"
	"encoding/hex"
	"hash/fnv"
)

func GenerateRandomString(length uint32) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomBytes), nil
}

// Uses the FNV-1a hash algorithm to calculate a uint64 from a generic string
func HashStringToUint64(str string) uint64 {
	h := fnv.New64a()

	h.Write([]byte(str))

	return h.Sum64()
}