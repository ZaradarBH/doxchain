package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceCodeRegistryKeyPrefix is the prefix to retrieve all DeviceCodeRegistry
	DeviceCodeRegistryKeyPrefix = "DeviceCodeRegistry/value/"
)

// DeviceCodeRegistryKey returns the store key to retrieve a DeviceCodeRegistry from the index fields
func DeviceCodeRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
