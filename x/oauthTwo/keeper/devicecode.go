package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/utils"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) DeviceCode(ctx sdk.Context, msg types.MsgDeviceCodeRequest) (response types.MsgDeviceCodeResponse, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	
	if err != nil {
		return response, err
	}
	
	isAuthorized, err := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, msg.TenantW3CIdentifier)

	if !isAuthorized {
		return response, err
	}

	tenantConfiguration, err := k.idpKeeper.GetTenantConfiguration(ctx, msg.TenantW3CIdentifier)

	if err != nil {
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

	response.DeviceCode, _ = utils.GenerateRandomString(32)
	response.UserCode, _ = utils.GenerateRandomString(8)
	response.VerificationUri = tenantConfiguration.LoginEndpoint

	tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, msg.TenantW3CIdentifier)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}
	
	ownerDid := didUtils.NewDidTokenFactory().Create(msg.Creator, "")
	deviceCodeRegistryEntry := idpTypes.DeviceCodeRegistryEntry{
		Owner: *ownerDid,
		DeviceCode: response.DeviceCode,
		UserCode:   response.UserCode,
		ExpiresAt:  ctx.BlockTime().Add(time.Minute * 15).Unix(),
	}

	tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes, deviceCodeRegistryEntry)

	k.idpKeeper.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

	return response, nil
}
