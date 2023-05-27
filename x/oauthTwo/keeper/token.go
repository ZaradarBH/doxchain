package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) Token(ctx sdk.Context, creator string, tenantW3CIdentifier string, clientRegistrationAppIdW3CIdentifier string, scope []string, clientSecret string, authorizationCode string, deviceCode string, clientAssertion string, clientAssertionType string, grantType types.GrantType) (accessToken string, tokenType string, expiresIn int64, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)

	if err != nil {
		return accessToken, tokenType, expiresIn, err
	}

	isAuthorized := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, tenantW3CIdentifier)

	if !isAuthorized {
		return accessToken, tokenType, expiresIn, err
	}

	var validScopes []string

	for _, requestedScope := range scope {
		validScopes = append(validScopes, k.idpKeeper.AuthorizeScope(ctx, tenantW3CIdentifier, clientRegistrationAppIdW3CIdentifier, requestedScope))
	}

	if len(validScopes) == 0 {
		return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	switch grantType {
	case types.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS_GRANT:
		return k.GenerateClientCredentialToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, clientSecret, clientAssertion, clientAssertionType)
	case types.GrantType_GRANT_TYPE_DEVICE_CODE_GRANT:
		return k.GenerateDeviceCodeToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, deviceCode)
	case types.GrantType_GRANT_TYPE_AUTHORIZATION_CODE_GRANT:
		return k.GenerateAuthorizationCodeToken(ctx, creator, tenantW3CIdentifier, scope, clientRegistrationAppIdW3CIdentifier, authorizationCode)
	}

	return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}
