package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
)

// Token method for simple oauth keeper
func (k Keeper) Token(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	switch msg.GrantType {
	case types.ClientCredentialsGrant.String():
		return k.GenerateClientCredentialToken(ctx, msg)		
	case types.DeviceCodeGrant.String():
		return k.GenerateDeviceCodeToken(ctx, msg)
	}

	return types.MsgTokenResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}

// AuthorizeRequest method for simple oauth keeper
func (k Keeper) AuthorizeRequest(ctx sdk.Context, msg types.MsgTokenRequest) (bool, error) {
	acl, err := k.idpKeeper.GetAccessClientList(ctx, msg.Tenant)

	if err != nil {
		return false, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.Creator == msg.Creator {
			return true, nil
		}
	}

	return false, nil
}
