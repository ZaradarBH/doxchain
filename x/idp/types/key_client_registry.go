package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClientRegistrationRegistryKeyPrefix is the prefix to retrieve all ClientRegistrationRegistry
	ClientRegistrationRegistryKeyPrefix = "ClientRegistrationRegistry/value/"
)

// ClientRegistrationRegistryKey returns the store key to retrieve a ClientRegistrationRegistry from the index fields
func ClientRegistrationRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
