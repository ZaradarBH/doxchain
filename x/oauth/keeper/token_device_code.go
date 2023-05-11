package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/dgrijalva/jwt-go"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	tenantDeviceCodes, found := k.GetDeviceCodes(ctx, msg.Tenant)

	if !found {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodes could not be found for tenant")
	}

	for index, deviceCodeEntry := range tenantDeviceCodes.Entries {
		if deviceCodeEntry.DeviceCode == msg.DeviceCode {
			jwtToken := jwt.New(jwt.SigningMethodHS256)
			claims := jwtToken.Claims.(jwt.MapClaims)

			claims["iss"] = types.ModuleName
			claims["sub"] = msg.ClientId
			//TODO: Implement oracle logic for adding unix timestamps to each block so we can use those to improve precision when issuing claims, assertions, etc
			claims["exp"] = time.Unix(int64(ctx.BlockHeight()), 0)

			signedToken, err := jwtToken.SignedString([]byte(msg.DeviceCode))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
			}

			//TODO: Save signed token to store until it is redeemed via /authorize
			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = "Bearer"

			//TODO: Make expire time configurable
			tokenResponse.ExpiresIn = 1800

			tenantDeviceCodes.Entries = append(tenantDeviceCodes.Entries[:index], tenantDeviceCodes.Entries[index+1:]...)
			
			k.SetDeviceCodes(ctx, tenantDeviceCodes)

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
