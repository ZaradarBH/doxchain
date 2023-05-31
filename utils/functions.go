package utils

import (
	"strconv"
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

func HashStringToUint64(str string) uint64 {
	h := fnv.New64a()

	h.Write([]byte(str))

	return h.Sum64()
}


func GetKeyBytes(indexer string) []byte {
	var key []byte

	//Ensure that indexers never exceed 20 chars
	if len(indexer) > 20 {
		indexer = strconv.FormatUint(HashStringToUint64(indexer), 10)
	}

	key = append(key, []byte(indexer)...)
	key = append(key, []byte("/")...)

	return key
}