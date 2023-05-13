package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/dgrijalva/jwt-go"
)

// Login method for simple idp keeper
func (k Keeper) Login(ctx sdk.Context, msg types.MsgAuthenticationRequest) (types.MsgAuthenticationResponse, error) {
	response := types.MsgAuthenticationResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IdpMasterKey)
	idpMasterKeyBytes := store.Get(byteKey)

	if idpMasterKeyBytes == nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not locate IDP master key in store")
	}

	//TODO: Move JwtTokenFactory to common util namespace, implement option for passing in claims in factory and replace this logic
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": msg.Creator,
		"exp": ctx.BlockTime().Add(time.Hour * 1),
	})

	tokenString, err := jwtToken.SignedString([]byte(idpMasterKeyBytes))
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
