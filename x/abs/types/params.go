package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
	
	didTypes "github.com/be-heroes/doxchain/x/did/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter store key
var (
	DefaultOperators = []didTypes.Did(nil) // no operators allowed
	DefaultBlockExpireOffset = sdk.NewUint(100000)

	ParamStoreKeyOperators = []byte("Operators")
	ParamStoreKeyBlockExpireOffset = []byte("BlockExpireOffset")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(operators []didTypes.Did, blockExpireOffset sdk.Uint) Params {
	return Params{
		Operators: operators,
		BlockExpireOffset: blockExpireOffset.Uint64(),
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultOperators, DefaultBlockExpireOffset)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyOperators, &p.Operators, validateDid),
		paramtypes.NewParamSetPair(ParamStoreKeyBlockExpireOffset, &p.BlockExpireOffset, validateBlockExpireOffset),
	}
}

func validateDid(i interface{}) error {
	_, ok := i.(didTypes.Did)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBlockExpireOffset(i interface{}) error {
	_, ok := i.(sdk.Uint)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	for _, operatorDid := range p.Operators {
		if err := validateDid(operatorDid); err != nil {
			return err
		}
	}

	if err := validateBlockExpireOffset(p.BlockExpireOffset); err != nil {
		return err
	}	

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}