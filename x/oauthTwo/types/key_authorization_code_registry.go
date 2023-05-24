package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	AuthorizationCodeRegistryKeyPrefix = "authorizationcoderegistry/value/"
)

func AuthorizationCodeRegistryKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
