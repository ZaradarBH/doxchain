package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	DeviceCodeRegistryKeyPrefix = "devicecoderegistry/value/"
)

func DeviceCodeRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
