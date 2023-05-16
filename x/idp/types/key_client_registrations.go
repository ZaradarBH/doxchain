package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClientRegistryKeyPrefix is the prefix to retrieve all ClientRegistry
	ClientRegistryKeyPrefix = "ClientRegistry/value/"
)

// ClientRegistryKey returns the store key to retrieve a ClientRegistry from the index fields
func ClientRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
