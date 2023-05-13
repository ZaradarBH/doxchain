package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
)

// DeviceCode method for simple oauth keeper
func (k Keeper) DeviceCode(ctx sdk.Context, msg types.MsgDeviceCodeRequest) (types.MsgDeviceCodeResponse, error) {
	//TODO: Implemeent devicecode logic

	return types.MsgDeviceCodeResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unknown")
}
