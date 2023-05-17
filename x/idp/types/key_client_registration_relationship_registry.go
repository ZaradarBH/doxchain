package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClientRegistrationRelationshipRegistryKeyPrefix is the prefix to retrieve all ClientRegistrationRelationship
	ClientRegistrationRelationshipRegistryKeyPrefix = "ClientRegistrationRelationshipRegistry/value/"
)

// ClientRegistrationRelationshipRegistryKey returns the store key to retrieve a ClientRegistrationRelationship from the index fields
func ClientRegistrationRelationshipRegistryKey(
	creator string,
) []byte {
	var key []byte

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
