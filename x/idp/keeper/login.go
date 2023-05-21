package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	"github.com/be-heroes/doxchain/x/idp/types"
)

func (k Keeper) Login(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) (signedToken string, err error) {
	isAuthorized, err := k.AuthorizeUser(ctx, user, tenantW3CIdentifier)

	if !isAuthorized {
		return "", err
	}

	jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(tenantW3CIdentifier, user.String(), user.String(), time.Minute*60)
	tokenString, err := jwtToken.SignedString([]byte(user.String()))

	if err != nil {
		return "", sdkerrors.Wrap(types.LoginError, "Could not issue refresh token")
	}

	return tokenString, nil
}

func (k Keeper) AuthorizeUser(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) (bool, error) {
	acl, err := k.GetAccessClientList(ctx, tenantW3CIdentifier)

	if err != nil {
		return false, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.User.Creator == user.String() {
			return true, nil
		}
	}

	return false, nil
}

func (k Keeper) AuthorizeScope(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationW3CIdentitifer string, scope string) (string, error) {
	registration, found := k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, clientRegistrationW3CIdentitifer)

	if !found {
		return "", sdkerrors.Wrap(types.LoginError, "Could not find a client registration matching the provided clientId")
	}

	for _, appScope := range registration.AppScopes {
		if appScope == scope {
			return appScope, nil
		}
	}

	return scope, sdkerrors.Wrap(types.LoginError, "Could not match any valid scopes")
}
