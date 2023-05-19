package types

import (
	"fmt"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceCodeRegistries: []DeviceCodeRegistry{},
		ClientRegistrationRegistries:     []ClientRegistrationRegistry{},
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	DeviceCodeRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceCodeRegistries {
		creator := string(DeviceCodeRegistryKey(elem.Owner.Creator))

		if _, ok := DeviceCodeRegistryIndexMap[creator]; ok {
			return fmt.Errorf("duplicated creator for DeviceCodeRegistry")
		}

		DeviceCodeRegistryIndexMap[creator] = struct{}{}
	}

	ClientRegistrationRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClientRegistrationRegistries {
		creator := string(ClientRegistrationRegistryKey(elem.Owner.Creator))
		
		if _, ok := ClientRegistrationRegistryIndexMap[creator]; ok {
			return fmt.Errorf("duplicated creator for ClientRegistrationRegistry")
		}
		
		ClientRegistrationRegistryIndexMap[creator] = struct{}{}
	}

	return gs.Params.Validate()
}
