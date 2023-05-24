package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetDotWellKnown(goCtx context.Context, req *types.QueryGetDotWellKnownRequest) (result *types.QueryGetDotWellKnownResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	configuration := k.GetTenantConfiguration(sdk.UnwrapSDKContext(goCtx), req.TenantW3CIdentifier)

	result.Issuer = configuration.Issuer
	result.AuthorizationEndpoint = configuration.AuthorizationEndpoint
	result.TokenEndpoint = configuration.TokenEndpoint
	result.TokenEndpointAuthMethodsSupported = configuration.TokenEndpointAuthMethodsSupported
	result.TokenEndpointAuthSigningAlgosSupported = configuration.TokenEndpointAuthSigningAlgosSupported
	result.UserInfoEndpoint = configuration.UserInfoEndpoint
	result.JwksUri = configuration.JwksUri
	result.RegistrationEndpoint = configuration.RegistrationEndpoint
	result.ScopesSupported = configuration.ScopesSupported
	result.ResponseTypesSupported = configuration.ResponseTypesSupported
	result.ServiceDocumentation = configuration.ServiceDocumentation
	result.UiLocalesSupported = configuration.UiLocalesSupported

	return result, nil
}
