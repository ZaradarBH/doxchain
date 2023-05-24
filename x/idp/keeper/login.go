package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	jwtUtils "github.com/be-heroes/doxchain/utils/jwt"
)

func (k Keeper) Login(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) string {
	isAuthorized := k.AuthorizeUser(ctx, user, tenantW3CIdentifier)

	if !isAuthorized {
		return ""
	}

	jwtToken := jwtUtils.NewJwtTokenFactory(jwtUtils.WithContext(&ctx)).Create(tenantW3CIdentifier, user.String(), user.String(), time.Minute*60)
	tokenString, err := jwtToken.SignedString([]byte(user.String()))

	if err != nil {
		return ""
	}

	return tokenString
}

func (k Keeper) AuthorizeUser(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) bool {
	acl := k.GetAccessClientList(ctx, tenantW3CIdentifier)

	for _, aclEntry := range acl.Entries {
		if aclEntry.User.Creator == user.String() {
			return true
		}
	}

	return false
}

func (k Keeper) AuthorizeScope(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationW3CIdentitifer string, scope string) string {
	registration, found := k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, clientRegistrationW3CIdentitifer)

	if !found {
		return ""
	}

	for _, appScope := range registration.AppScopes {
		if appScope == scope {
			return appScope
		}
	}

	return ""
}
