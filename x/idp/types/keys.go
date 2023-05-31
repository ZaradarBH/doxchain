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
	TenantRegistryKeyPrefix = "tenantregistry/value/"
	ClientRegistrationRegistryKeyPrefix = "clientregistrationregistry/value/"
	ClientRegistrationRelationshipRegistryKeyPrefix = "clientregistrationrelationshipregistry/value/"
	DeviceCodeRegistryKeyPrefix = "devicecoderegistry/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
