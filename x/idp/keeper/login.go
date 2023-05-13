package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/be-heroes/doxchain/utils"
	"github.com/dgrijalva/jwt-go"
)

// Login method for simple idp keeper
func (k Keeper) Login(ctx sdk.Context, msg types.MsgAuthenticationRequest) (types.MsgAuthenticationResponse, error) {
	response := types.MsgAuthenticationResponse{}
	isAuthorized, err := k.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}
	
	jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.Tenant, msg.Creator, msg.ClientId, time.Minute * 60)
	tokenString, err := jwtToken.SignedString([]byte(msg.Creator))

	if err != nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not issue refresh token")
	}

	response.Token = tokenString

	return response, nil
}

// AuthorizeCreator checks if a creator belongs to a given tenant
func (k Keeper) AuthorizeCreator(ctx sdk.Context, tenant string, creator string) (bool, error) {
	acl, err := k.GetAccessClientList(ctx, tenant)

	if err != nil {
		return false, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.Creator == creator {
			return true, nil
		}
	}

	return false, nil
}
