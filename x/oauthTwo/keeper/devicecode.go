package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/be-heroes/doxchain/utils"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

// DeviceCode method for simple oauth keeper
func (k Keeper) DeviceCode(ctx sdk.Context, msg types.MsgDeviceCodeRequest) (types.MsgDeviceCodeResponse, error) {
	response := types.MsgDeviceCodeResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	tenantConfiguration, err := k.idpKeeper.GetTenantConfiguration(ctx, msg.Tenant)

	if err != nil {
		return response, err
	}

	//TODO: Validate ClientId and Scope
	response.DeviceCode, _ = utils.GenerateRandomString(32)
	response.UserCode, _ = utils.GenerateRandomString(8)
	response.VerificationUri = tenantConfiguration.AuthorizationEndpoint

	tenantDeviceCodeRegistry, found := k.GetDeviceCodeRegistry(ctx, msg.Tenant)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}

	deviceCodeEntry := types.DeviceCodeEntry{
		Creator:    msg.Creator,
		DeviceCode: response.DeviceCode,
		ExpiresAt:  ctx.BlockTime().Add(time.Minute * 15).Unix(),
	}

	tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes, deviceCodeEntry)

	k.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

	return response, nil
}
