package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceCodeRegistries: []DeviceCodeRegistry{},
		ClientRegistries:     []ClientRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in DeviceCodeRegistry
	DeviceCodeRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceCodeRegistries {
		creator := string(DeviceCodeRegistryKey(elem.Owner.Creator))
		if _, ok := DeviceCodeRegistryIndexMap[creator]; ok {
			return fmt.Errorf("duplicated creator for DeviceCodeRegistry")
		}
		DeviceCodeRegistryIndexMap[creator] = struct{}{}
	}

	// Check for duplicated index in ClientRegistry
	ClientRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClientRegistries {
		creator := string(ClientRegistryKey(elem.Owner.Creator))
		if _, ok := ClientRegistryIndexMap[creator]; ok {
			return fmt.Errorf("duplicated creator for ClientRegistry")
		}
		ClientRegistryIndexMap[creator] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
