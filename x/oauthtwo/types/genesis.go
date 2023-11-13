package types

import (
	"fmt"
	utils "github.com/be-heroes/doxchain/utils"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccessTokenRegistries:       []AccessTokenRegistry{},
		AuthorizationCodeRegistries: []AuthorizationCodeRegistry{},
		Params:                      DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	AccessTokenRegistryIndexMap := make(map[string]struct{})

	for _, registry := range gs.AccessTokenRegistries {
		indexer := string(utils.GetKeyBytes(registry.Owner.Creator))

		if _, ok := AccessTokenRegistryIndexMap[indexer]; ok {
			return fmt.Errorf("duplicated indexer for AccessTokenRegistry")
		}

		AccessTokenRegistryIndexMap[indexer] = struct{}{}
	}

	authorizationCodeRegistryIndexMap := make(map[string]struct{})

	for _, registry := range gs.AuthorizationCodeRegistries {
		indexer := string(utils.GetKeyBytes(registry.Owner.Creator))

		if _, ok := authorizationCodeRegistryIndexMap[indexer]; ok {
			return fmt.Errorf("duplicated indexer for authorizationCodeRegistry")
		}

		authorizationCodeRegistryIndexMap[indexer] = struct{}{}
	}

	return gs.Params.Validate()
}
