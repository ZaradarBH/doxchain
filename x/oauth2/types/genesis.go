package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceCodeRegistryList:  []DeviceCodeRegistry{},
		AccessTokenRegistryList: []AccessTokenRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in DeviceCodeRegistry
	DeviceCodeRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceCodeRegistryList {
		tenant := string(DeviceCodeRegistryKey(elem.Tenant))
		if _, ok := DeviceCodeRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated index for DeviceCodeRegistry")
		}
		DeviceCodeRegistryIndexMap[tenant] = struct{}{}
	}
	// Check for duplicated index in AccessTokenRegistry
	AccessTokenRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccessTokenRegistryList {
		tenant := string(AccessTokenRegistryKey(elem.Tenant))
		if _, ok := AccessTokenRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated index for AccessTokenRegistry")
		}
		AccessTokenRegistryIndexMap[tenant] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
