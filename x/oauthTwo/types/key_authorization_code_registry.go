package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AuthorizationCodeRegistryKeyPrefix is the prefix to retrieve all AuthorizationCodeRegistry
	AuthorizationCodeRegistryKeyPrefix = "AuthorizationCodeRegistry/value/"
)

// AuthorizationCodeRegistryKey returns the store key to retrieve a AuthorizationCodeRegistry from the index fields
func AuthorizationCodeRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
