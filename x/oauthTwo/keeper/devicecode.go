package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/utils"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) DeviceCode(ctx sdk.Context, creator string, tenantW3CIdentifier string, clientRegistrationAppIdW3CIdentifier string, scope []string) (deviceCode string, userCode string, verificationUri string, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)

	if err != nil {
		return deviceCode, userCode, verificationUri, err
	}

	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "DeviceCode", creator)

	if err != nil {
		return deviceCode, userCode, verificationUri, err
	}

	isAuthorized, err := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, tenantW3CIdentifier)

	if !isAuthorized {
		return deviceCode, userCode, verificationUri, err
	}

	tenantConfiguration, err := k.idpKeeper.GetTenantConfiguration(ctx, tenantW3CIdentifier)

	if err != nil {
		return deviceCode, userCode, verificationUri, err
	}

	var validScopes []string

	for _, requestedScope := range scope {
		validScope, err := k.idpKeeper.AuthorizeScope(ctx, tenantW3CIdentifier, clientRegistrationAppIdW3CIdentifier, requestedScope)

		if err != nil {
			return deviceCode, userCode, verificationUri, err
		}

		validScopes = append(validScopes, validScope)
	}

	if len(validScopes) == 0 {
		return deviceCode, userCode, verificationUri, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	deviceCode, _ = utils.GenerateRandomString(32)
	userCode, _ = utils.GenerateRandomString(8)
	verificationUri = tenantConfiguration.LoginEndpoint

	tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, tenantW3CIdentifier)

	if !found {
		return deviceCode, userCode, verificationUri, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}

	ownerDid := didUtils.NewDidTokenFactory().Create(creator, didUrl)
	deviceCodeRegistryEntry := idpTypes.DeviceCodeRegistryEntry{
		Owner:      *ownerDid,
		DeviceCode: deviceCode,
		UserCode:   userCode,
		ExpiresAt:  ctx.BlockTime().Add(time.Minute * 15).Unix(),
	}

	tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes, deviceCodeRegistryEntry)

	k.idpKeeper.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

	return deviceCode, userCode, verificationUri, nil
}
