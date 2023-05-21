package keeper

import (
	"fmt"
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SetClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationRelationshipRegistryEntry types.ClientRegistrationRelationshipRegistryEntry) error {
	ownerRegistration, found := k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier())

	if !found {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid owner")
	}

	destinationRegistration, found := k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier())

	if !found {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid destination")
	}

	for _, aclEntry := range clientRegistrationRelationshipRegistryEntry.AccessClientList.Entries {
		matchOwner := false;
		matchDestination := false;

		for _, aclEntryOwner := range ownerRegistration.AccessClientList.Entries {
			if aclEntryOwner.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchOwner = true;
			}
		}

		for _, aclEntryDestination := range destinationRegistration.AccessClientList.Entries {
			if aclEntryDestination.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchDestination = true;
			}
		}

		if !matchOwner || !matchDestination {
			return sdkerrors.Wrap(types.AccessClientListError, "Illigal relationship. Acl must match both owner and destination acls")
		}
	}

	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))
	
	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		//TODO: Put the @ seperator in a constant somewhere
		fmt.Sprintf("%s@%s", clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier(), clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier()),
	))
		
	if b == nil {
		clientRegistrationRelationshipRegistry = types.ClientRegistrationRelationshipRegistry{}
	} else {
		k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)	
	}
	
	for index, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier() == existingEntry.Owner.GetW3CIdentifier() && clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier() == existingEntry.Destination.GetW3CIdentifier() {
			clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships[:index], clientRegistrationRelationshipRegistry.Relationships[index+1:]...)
		}
	}

	clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships, clientRegistrationRelationshipRegistryEntry)

	b = k.cdc.MustMarshal(&clientRegistrationRelationshipRegistry)

	store.Set(types.ClientRegistrationRelationshipRegistryKey(
		//TODO: Put the @ seperator in a constant somewhere
		fmt.Sprintf("%s@%s", clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier(), clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier()),
	), b)

	return nil
}

func (k Keeper) RemoveClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, ownerClientRegistrationW3CIdentitifer string, destinationClientRegistrationW3CIdentitifer string) error {
	_, found := k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, ownerClientRegistrationW3CIdentitifer)

	if !found {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid owner")
	}

	_, found = k.GetClientRegistration(ctx, clientRegistrationRegistryW3CIdentitifer, destinationClientRegistrationW3CIdentitifer)

	if !found {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid destination")
	}
	
	//TODO: Finish this method

	return nil
}