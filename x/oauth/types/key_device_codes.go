package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceCodesKeyPrefix is the prefix to retrieve all DeviceCodes
	DeviceCodesKeyPrefix = "DeviceCodes/value/"
)

// DeviceCodesKey returns the store key to retrieve a DeviceCodes from the index fields
func DeviceCodesKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
