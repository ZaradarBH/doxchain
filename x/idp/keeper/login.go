package keeper

import (
	"time"
	"github.com/decred/dcrd/dcrec/secp256k1"
    b64 "encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	jwtUtils "github.com/be-heroes/doxchain/utils/jwt"
)

func (k Keeper) Login(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) string {
	isAuthorized := k.AuthorizeUser(ctx, user, tenantW3CIdentifier)

	if !isAuthorized {
		return ""
	}

	account := k.accountKeeper.GetAccount(ctx, user)

	if account == nil {
		return ""
	}

	pubKey, err := secp256k1.ParsePubKey(account.GetPubKey().Bytes())

	if err != nil {
		return ""
	}

	jwtToken := jwtUtils.NewJwtTokenFactory(jwtUtils.WithContext(&ctx)).Create(tenantW3CIdentifier, user.String(), user.String(), time.Minute*60)
	tokenString, err := jwtToken.SignedString(pubKey)

	if err != nil {
		return ""
	}
	
	ciphertext, err := secp256k1.Encrypt(pubKey, []byte(tokenString))

	if err != nil {
		return ""
	}

	return b64.StdEncoding.EncodeToString(ciphertext)
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
