package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) Token(ctx sdk.Context, creator string, tenantW3CIdentifier string, clientRegistrationAppIdW3CIdentifier string, scope []string, clientSecret string, authorizationCode string, deviceCode string, clientAssertion string, clientAssertionType string, grantType string) (accessToken string, tokenType string, expiresIn int64, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	
	if err != nil {
		return accessToken, tokenType, expiresIn, err
	}

	isAuthorized, err := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, tenantW3CIdentifier)

	if !isAuthorized {
		return accessToken, tokenType, expiresIn, err
	}
	
	var validScopes []string

	for _, requestedScope := range scope {
		validScope, err := k.idpKeeper.AuthorizeScope(ctx, tenantW3CIdentifier, clientRegistrationAppIdW3CIdentifier, requestedScope)

		if err != nil {
			return accessToken, tokenType, expiresIn, err
		}

		validScopes = append(validScopes, validScope)
	}

	if len(validScopes) == 0 {
		return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	switch grantType {
	case types.ClientCredentialsGrant.String():
		return k.GenerateClientCredentialToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, clientSecret, clientAssertion, clientAssertionType)
	case types.DeviceCodeGrant.String():
		return k.GenerateDeviceCodeToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, deviceCode)
	case types.AuthorizationCodeGrant.String():
		return k.GenerateAuthorizationCodeToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, authorizationCode)
	}

	return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}
