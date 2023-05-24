package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	ClientRegistrationRelationshipRegistryKeyPrefix = "ClientRegistrationRelationshipRegistry/value/"
)

func ClientRegistrationRelationshipRegistryKey(
	creator string,
) []byte {
	var key []byte

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
