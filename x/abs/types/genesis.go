package types

import (
"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PartitionedPoolRegistryList: []PartitionedPoolRegistry{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in partitionedPools
partitionedPoolsIndexMap := make(map[string]struct{})

for _, elem := range gs.PartitionedPoolRegistryList {
	creator := string(PartitionedPoolRegistryKey(elem.Creator))
	if _, ok := partitionedPoolsIndexMap[creator]; ok {
		return fmt.Errorf("duplicated creator for PartitionedPoolRegistries")
	}
	partitionedPoolsIndexMap[creator] = struct{}{}
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
