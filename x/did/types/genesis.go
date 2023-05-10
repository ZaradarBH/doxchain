package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidList: []Did{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in did
	didIdMap := make(map[string]bool)
	for _, elem := range gs.DidList {
		fullyQualifiedDidIdentifier := fmt.Sprintf("did:%s:%s", elem.MethodName, elem.MethodId)

		if _, ok := didIdMap[fullyQualifiedDidIdentifier]; ok {
			return fmt.Errorf("duplicated id for did")
		}		
		didIdMap[fullyQualifiedDidIdentifier] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
