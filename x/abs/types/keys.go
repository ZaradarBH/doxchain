package types

const (
	// ModuleName defines the module name
	ModuleName = "abs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_abs"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// Keys for abs store
// Items are stored with the following key: values
//
// - 0x01: sdk.Dec
// - 0x02: []WatchlistEntry
var (
	// Keys for store prefixed
	BreakFactorKey = []byte{0x01}
	WatchlistKey   = []byte{0x02}
)
