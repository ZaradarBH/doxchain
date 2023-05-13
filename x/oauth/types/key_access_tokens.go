package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccessTokensKeyPrefix is the prefix to retrieve all AccessTokens
	AccessTokensKeyPrefix = "AccessTokens/value/"
)

// AccessTokensKey returns the store key to retrieve a AccessTokens from the index fields
func AccessTokensKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
