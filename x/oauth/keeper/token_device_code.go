package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/utils"
	"github.com/be-heroes/doxchain/x/oauth/types"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	tenantDeviceCodes, found := k.GetDeviceCodes(ctx, msg.Tenant)

	if !found {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodes could not be found for tenant")
	}

	for index, deviceCodeEntry := range tenantDeviceCodes.Entries {
		if deviceCodeEntry.DeviceCode == msg.DeviceCode {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(&msg)
			signedToken, err := jwtToken.SignedString([]byte(msg.DeviceCode))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
			}

			//TODO: Save signed token to store until it is removed, if we even want to do that?
			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = types.Bearer.String()

			//TODO: Make expire time configurable!
			tokenResponse.ExpiresIn = 1800

			tenantDeviceCodes.Entries = append(tenantDeviceCodes.Entries[:index], tenantDeviceCodes.Entries[index+1:]...)
			
			k.SetDeviceCodes(ctx, tenantDeviceCodes)

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
