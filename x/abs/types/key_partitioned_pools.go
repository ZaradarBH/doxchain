package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // PartitionedPoolsKeyPrefix is the prefix to retrieve all PartitionedPools
	PartitionedPoolsKeyPrefix = "PartitionedPools/value/"
)

// PartitionedPoolsKey returns the store key to retrieve a PartitionedPools from the index fields
func PartitionedPoolsKey(
index string,
) []byte {
	var key []byte
    
    indexBytes := []byte(index)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}