package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetDotWellKnown(goCtx context.Context, req *types.QueryGetDotWellKnownRequest) (*types.QueryGetDotWellKnownResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	configuration, err := k.GetTenantConfiguration(sdk.UnwrapSDKContext(goCtx), req.Tenant)

	if err != nil {
		return nil, err
	}

	response := &types.QueryGetDotWellKnownResponse{
		Issuer: configuration.Issuer,
		AuthorizationEndpoint: configuration.AuthorizationEndpoint,
		TokenEndpoint: configuration.TokenEndpoint,
		TokenEndpointAuthMethodsSupported: configuration.TokenEndpointAuthMethodsSupported,
		TokenEndpointAuthSigningAlgosSupported: configuration.TokenEndpointAuthSigningAlgosSupported,
		UserInfoEndpoint: configuration.UserInfoEndpoint,
		JwksUri: configuration.JwksUri,
		RegistrationEndpoint: configuration.RegistrationEndpoint,
		ScopesSupported: configuration.ScopesSupported,
		ResponseTypesSupported: configuration.ResponseTypesSupported,
		ServiceDocumentation: configuration.ServiceDocumentation,
		UiLocalesSupported: configuration.UiLocalesSupported,
	}
	
	return response, nil
}
