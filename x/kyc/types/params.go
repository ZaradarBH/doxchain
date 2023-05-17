package types

import (
	"fmt"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter store key
var (
	DefaultApprovers = []didTypes.Did(nil) // none allowed

	ParamStoreKeyApprovers = []byte("Approvers")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(approvers []didTypes.Did) Params {
	return Params{
		Approvers: approvers,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultApprovers)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyApprovers, &p.Approvers, validateApprovers),
	}
}

func validateApprovers(i interface{}) error {
	_, ok := i.([]didTypes.Did)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateApprovers(p.Approvers); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
