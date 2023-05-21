package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetClientRegistration(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistration types.ClientRegistration) error {
	clientRegistry, found := k.GetClientRegistrationRegistry(ctx, clientRegistrationRegistryW3CIdentitifer)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetW3CIdentifier() == clientRegistration.Id.GetW3CIdentifier() {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				break
			}
		}

		clientRegistry.Registrations = append(clientRegistry.Registrations, clientRegistration)

		k.SetClientRegistrationRegistry(ctx, clientRegistry)
	}

	return nil
}

func (k Keeper) RemoveClientRegistration(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationW3CIdentitifer string) {
	//TODO: Check if clientRegistrationW3CIdentitifer is a well-formed did string.
	clientRegistry, found := k.GetClientRegistrationRegistry(ctx, clientRegistrationRegistryW3CIdentitifer)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetW3CIdentifier() == clientRegistrationW3CIdentitifer {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				k.SetClientRegistrationRegistry(ctx, clientRegistry)

				break
			}
		}
	}
}

func (k Keeper) GetClientRegistration(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationW3CIdentitifer string) (val types.ClientRegistration, found bool) {	
	//TODO: Check if clientRegistrationW3CIdentitifer is a well-formed did string.
	clientRegistry, found := k.GetClientRegistrationRegistry(ctx, clientRegistrationRegistryW3CIdentitifer)

	if found {
		for _, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetW3CIdentifier() == clientRegistrationW3CIdentitifer {
				return existingClientRegistration, found
			}
		}
	}
	
	return val, found
}
