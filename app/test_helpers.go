package app

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
)

var defaultGenesisBz []byte

func getDefaultGenesisStateBytes(cdc codec.JSONCodec) []byte {
	if len(defaultGenesisBz) == 0 {
		genesisState := NewDefaultGenesisState(cdc)
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}
		defaultGenesisBz = stateBytes
	}
	return defaultGenesisBz
}

// Setup initializes a new OsmosisApp.
func Setup(isCheckTx bool) *App {
	db := dbm.NewMemDB()
	encoding := MakeEncodingConfig()

	app := New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encoding, simapp.EmptyAppOptions{})
	if !isCheckTx {
		stateBytes := getDefaultGenesisStateBytes(encoding.Marshaler)

		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}
