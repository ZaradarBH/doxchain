package types

import (
	"fmt"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultApprovers = []didTypes.Did(nil) // none allowed

	ParamStoreKeyApprovers = []byte("Approvers")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(approvers []didTypes.Did) Params {
	return Params{
		Approvers: approvers,
	}
}

func DefaultParams() Params {
	return NewParams(DefaultApprovers)
}

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

func (p Params) Validate() error {
	if err := validateApprovers(p.Approvers); err != nil {
		return err
	}

	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
