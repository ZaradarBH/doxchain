package types

import (
	"fmt"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccessTokenRegistries:       []AccessTokenRegistry{},
		AuthorizationCodeRegistries: []AuthorizationCodeRegistry{},
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	AccessTokenRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccessTokenRegistries {
		tenant := string(AccessTokenRegistryKey(elem.Owner.Creator))

		if _, ok := AccessTokenRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated tenant for AccessTokenRegistry")
		}
		
		AccessTokenRegistryIndexMap[tenant] = struct{}{}
	}
	
	authorizationCodeRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.AuthorizationCodeRegistries {
		tenant := string(AuthorizationCodeRegistryKey(elem.Owner.Creator))

		if _, ok := authorizationCodeRegistryIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated tenant for authorizationCodeRegistry")
		}

		authorizationCodeRegistryIndexMap[tenant] = struct{}{}
	}

	return gs.Params.Validate()
}
