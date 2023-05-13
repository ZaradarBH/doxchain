package keeper

import (
	"time"	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth2/types"
	"github.com/be-heroes/doxchain/utils"
)

// DeviceCode method for simple oauth keeper
func (k Keeper) DeviceCode(ctx sdk.Context, msg types.MsgDeviceCodeRequest) (types.MsgDeviceCodeResponse, error) {
	response := types.MsgDeviceCodeResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	response.DeviceCode, _ = utils.GenerateRandomString(32)
	response.UserCode, _ = utils.GenerateRandomString(8)
	//TODO: Implement support for verification uri in tenant
	response.VerificationUri = "http://tenant_verification_uri/"

	tenantDeviceCodeRegistry, found := k.GetDeviceCodeRegistry(ctx, msg.Tenant)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}
	
	deviceCodeInfo := types.DeviceCodeInfo{
		Creator: msg.Creator,
		DeviceCode: response.DeviceCode,
		ExpiresAt: ctx.BlockTime().Add(time.Minute * 15).Unix(),
	}

	tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes, deviceCodeInfo)

	k.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

	return response, nil
}
