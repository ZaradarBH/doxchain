package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) Token(ctx sdk.Context, msg types.MsgTokenRequest) (response types.MsgTokenResponse, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	
	if err != nil {
		return response, err
	}

	isAuthorized, err := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, msg.TenantW3CIdentifier)

	if !isAuthorized {
		return response, err
	}
	
	var validScopes []string

	for _, requestedScope := range msg.Scope {
		validScope, err := k.idpKeeper.AuthorizeScope(ctx, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, requestedScope)

		if err != nil {
			return response, err
		}

		validScopes = append(validScopes, validScope)
	}

	if len(validScopes) == 0 {
		return response, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	switch msg.GrantType {
	case types.ClientCredentialsGrant.String():
		return k.GenerateClientCredentialToken(ctx, msg)
	case types.DeviceCodeGrant.String():
		return k.GenerateDeviceCodeToken(ctx, msg)
	}

	return types.MsgTokenResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}
