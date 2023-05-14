package types

import (
"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PartitionedPoolsList: []PartitionedPools{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in partitionedPools
partitionedPoolsIndexMap := make(map[string]struct{})

for _, elem := range gs.PartitionedPoolsList {
	index := string(PartitionedPoolsKey(elem.Index))
	if _, ok := partitionedPoolsIndexMap[index]; ok {
		return fmt.Errorf("duplicated index for partitionedPools")
	}
	partitionedPoolsIndexMap[index] = struct{}{}
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
