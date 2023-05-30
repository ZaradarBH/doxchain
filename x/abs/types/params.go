package types

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	DefaultBlockExpireOffset = sdk.NewInt(100000)
	DefaultBreakFactor       = sdk.MustNewDecFromStr("0.5")

	ParamStoreKeyBlockExpireOffset = []byte("BlockExpireOffset")
	ParamStoreKeyBreakFactor       = []byte("BreakFactor")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(breakFactor sdk.Dec, blockExpireOffset math.Int) Params {
	return Params{
		BlockExpireOffset: blockExpireOffset,
		BreakFactor:       breakFactor,
	}
}

func DefaultParams() Params {
	return NewParams(DefaultBreakFactor, DefaultBlockExpireOffset)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyBlockExpireOffset, &p.BlockExpireOffset, validateBlockExpireOffset),
		paramtypes.NewParamSetPair(ParamStoreKeyBreakFactor, &p.BreakFactor, validateBreakFactor),
	}
}

func validateBlockExpireOffset(i interface{}) error {
	_, ok := i.(math.Int)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBreakFactor(i interface{}) error {
	p, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// break factor must be between 0 and 1
	if p.IsNegative() || p.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrBreakFactorOutOfBounds, p.String())
	}

	return nil
}

func (p Params) Validate() error {
	if err := validateBlockExpireOffset(p.BlockExpireOffset); err != nil {
		return err
	}

	if err := validateBreakFactor(p.BreakFactor); err != nil {
		return err
	}

	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
