package types

const (
	// ModuleName defines the module name
	ModuleName = "idp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_idp"
)

const (
	// TenantRegistryKeyPrefix is the prefix to retrieve all TenantRegistry
	TenantRegistryKeyPrefix = "TenantRegistry/value/"
)

// TenantRegistryKey returns the store key to retrieve a TenantRegistryKey from the identifier fields
func TenantRegistryKey(
	identifier string,
) []byte {
	var key []byte

	identifierBytes := []byte(identifier)
	key = append(key, identifierBytes...)
	key = append(key, []byte("/")...)

	return key
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
