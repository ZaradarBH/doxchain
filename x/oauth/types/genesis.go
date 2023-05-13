package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DeviceCodesList:  []DeviceCodes{},
		AccessTokensList: []AccessTokens{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in deviceCodes
	deviceCodesIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceCodesList {
		tenant := string(DeviceCodesKey(elem.Tenant))
		if _, ok := deviceCodesIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated index for deviceCodes")
		}
		deviceCodesIndexMap[tenant] = struct{}{}
	}
	// Check for duplicated index in accessTokens
	accessTokensIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccessTokensList {
		tenant := string(AccessTokensKey(elem.Tenant))
		if _, ok := accessTokensIndexMap[tenant]; ok {
			return fmt.Errorf("duplicated index for accessTokens")
		}
		accessTokensIndexMap[tenant] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
