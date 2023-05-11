package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	acl, err := k.idpKeeper.GetAccessClientList(ctx, msg.Tenant)

	if err != nil {
		return tokenResponse, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.ClientId == msg.ClientId {
			//TODO: Implement devicecode / account map
			//TODO: Generate device code and update devicecodelist

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
