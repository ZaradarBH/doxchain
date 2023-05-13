package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccessTokenRegistryKeyPrefix is the prefix to retrieve all AccessTokenRegistry
	AccessTokenRegistryKeyPrefix = "AccessTokenRegistry/value/"
)

// AccessTokenRegistryKey returns the store key to retrieve a AccessTokenRegistry from the index fields
func AccessTokenRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
