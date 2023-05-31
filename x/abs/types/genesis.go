package types

import (
	"fmt"
	utils "github.com/be-heroes/doxchain/utils"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PartitionedPoolRegistries: []PartitionedPoolRegistry{},
		Params:                    DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	partitionedPoolsIndexMap := make(map[string]struct{})

	for _, registry := range gs.PartitionedPoolRegistries {
		indexer := string(utils.GetKeyBytes(registry.Owner.Creator))

		if _, ok := partitionedPoolsIndexMap[indexer]; ok {
			return fmt.Errorf("duplicated indexer for PartitionedPoolRegistries")
		}

		partitionedPoolsIndexMap[indexer] = struct{}{}
	}

	return gs.Params.Validate()
}
