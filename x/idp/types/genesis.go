package types

import (
	"fmt"
	utils "github.com/be-heroes/doxchain/utils"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceCodeRegistries:         []DeviceCodeRegistry{},
		ClientRegistrationRegistries: []ClientRegistrationRegistry{},
		Params:                       DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	DeviceCodeRegistryIndexMap := make(map[string]struct{})

	for _, registry := range gs.DeviceCodeRegistries {
		indexer := string(utils.GetKeyBytes(registry.Owner.Creator))

		if _, ok := DeviceCodeRegistryIndexMap[indexer]; ok {
			return fmt.Errorf("duplicated indexer for DeviceCodeRegistry")
		}

		DeviceCodeRegistryIndexMap[indexer] = struct{}{}
	}

	ClientRegistrationRegistryIndexMap := make(map[string]struct{})

	for _, registry := range gs.ClientRegistrationRegistries {
		indexer := string(utils.GetKeyBytes(registry.Owner.Creator))

		if _, ok := ClientRegistrationRegistryIndexMap[indexer]; ok {
			return fmt.Errorf("duplicated indexer for ClientRegistrationRegistry")
		}

		ClientRegistrationRegistryIndexMap[indexer] = struct{}{}
	}

	return gs.Params.Validate()
}
