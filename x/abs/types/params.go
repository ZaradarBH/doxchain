package types

import (
	"fmt"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultOperators         = []didTypes.Did(nil) // none allowed
	DefaultBlockExpireOffset = sdk.NewInt(100000)

	ParamStoreKeyOperators         = []byte("Operators")
	ParamStoreKeyBlockExpireOffset = []byte("BlockExpireOffset")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(operators []didTypes.Did, blockExpireOffset sdk.Int) Params {
	return Params{
		Operators:         operators,
		BlockExpireOffset: blockExpireOffset,
	}
}

func DefaultParams() Params {
	return NewParams(DefaultOperators, DefaultBlockExpireOffset)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyOperators, &p.Operators, validateOperators),
		paramtypes.NewParamSetPair(ParamStoreKeyBlockExpireOffset, &p.BlockExpireOffset, validateBlockExpireOffset),
	}
}

func validateBlockExpireOffset(i interface{}) error {
	_, ok := i.(sdk.Int)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateOperators(i interface{}) error {
	_, ok := i.([]didTypes.Did)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func (p Params) Validate() error {
	if err := validateBlockExpireOffset(p.BlockExpireOffset); err != nil {
		return err
	}

	if err := validateOperators(p.Operators); err != nil {
		return err
	}

	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
