package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	PartitionedPoolRegistryKeyPrefix = "PartitionedPoolRegistry/value/"
)

func PartitionedPoolRegistryKey(
	creator string,
) []byte {
	var key []byte

	indexBytes := []byte(creator)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
