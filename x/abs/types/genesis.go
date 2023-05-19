package types

import (
	"fmt"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PartitionedPoolRegistries: []PartitionedPoolRegistry{},
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	partitionedPoolsIndexMap := make(map[string]struct{})

	for _, elem := range gs.PartitionedPoolRegistries {
		creator := string(PartitionedPoolRegistryKey(elem.Owner.Creator))

		if _, ok := partitionedPoolsIndexMap[creator]; ok {
			return fmt.Errorf("duplicated creator for PartitionedPoolRegistries")
		}
		
		partitionedPoolsIndexMap[creator] = struct{}{}
	}

	return gs.Params.Validate()
}
