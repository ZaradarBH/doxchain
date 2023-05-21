package types

import (
	"fmt"
)

const DefaultIndex uint64 = 1

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidList: []Did{},
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	didIdMap := make(map[string]bool)

	for _, elem := range gs.DidList {
		didW3CIdentifier := fmt.Sprintf("did:%s:%s", elem.MethodName, elem.MethodId)

		if _, ok := didIdMap[didW3CIdentifier]; ok {
			return fmt.Errorf("duplicated id for did")
		}
		
		didIdMap[didW3CIdentifier] = true
	}

	return gs.Params.Validate()
}
