package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // PartitionedPoolRegistryKeyPrefix is the prefix to retrieve all PartitionedPools
	PartitionedPoolRegistryKeyPrefix = "PartitionedPoolRegistry/value/"
)

// PartitionedPoolsKey returns the store key to retrieve a PartitionedPools from the index fields
func PartitionedPoolRegistryKey(
creator string,
) []byte {
	var key []byte
    
    indexBytes := []byte(creator)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}