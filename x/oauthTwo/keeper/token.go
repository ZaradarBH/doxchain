package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

// Token method for simple oauth keeper
func (k Keeper) Token(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	//TODO: Validate ClientId and Scope
	switch msg.GrantType {
	case types.ClientCredentialsGrant.String():
		return k.GenerateClientCredentialToken(ctx, msg)
	case types.DeviceCodeGrant.String():
		return k.GenerateDeviceCodeToken(ctx, msg)
	}

	return types.MsgTokenResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}
