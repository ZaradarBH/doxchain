package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, ownerClientRegistrationW3CIdentitifer string, destinationClientRegistrationW3CIdentitifer string) (result types.ClientRegistrationRelationshipRegistryEntry, found bool) {
	for _, relationship := range k.GetClientRegistrationRelationships(ctx, clientRegistrationRegistryW3CIdentitifer, ownerClientRegistrationW3CIdentitifer) {
		if relationship.Destination.GetW3CIdentifier() == destinationClientRegistrationW3CIdentitifer {
			return relationship, true
		}
	}

	return result, false
}

func (k Keeper) GetClientRegistrationRelationships(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, ownerClientRegistrationW3CIdentitifer string) (result []types.ClientRegistrationRelationshipRegistryEntry) {
	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))

	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRegistryW3CIdentitifer,
	))

	k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)

	for _, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if existingEntry.Owner.GetW3CIdentifier() == ownerClientRegistrationW3CIdentitifer {
			result = append(result, existingEntry)
		}
	}

	return result
}

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
		matchOwner := false
		matchDestination := false

		for _, aclEntryOwner := range ownerRegistration.AccessClientList.Entries {
			if aclEntryOwner.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchOwner = true
			}
		}

		for _, aclEntryDestination := range destinationRegistration.AccessClientList.Entries {
			if aclEntryDestination.User.GetW3CIdentifier() == aclEntry.User.GetW3CIdentifier() {
				matchDestination = true
			}
		}

		if !matchOwner || !matchDestination {
			return sdkerrors.Wrap(types.AccessClientListError, "Illigal relationship. Acl must match both owner and destination acls")
		}
	}

	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))

	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRegistryW3CIdentitifer,
	))

	k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)

	for index, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if clientRegistrationRelationshipRegistryEntry.Owner.GetW3CIdentifier() == existingEntry.Owner.GetW3CIdentifier() && clientRegistrationRelationshipRegistryEntry.Destination.GetW3CIdentifier() == existingEntry.Destination.GetW3CIdentifier() {
			clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships[:index], clientRegistrationRelationshipRegistry.Relationships[index+1:]...)
		}
	}

	clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships, clientRegistrationRelationshipRegistryEntry)

	b = k.cdc.MustMarshal(&clientRegistrationRelationshipRegistry)

	store.Set(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRegistryW3CIdentitifer,
	), b)

	return nil
}

func (k Keeper) RemoveClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, ownerClientRegistrationW3CIdentitifer string, destinationClientRegistrationW3CIdentitifer string) error {
	var clientRegistrationRelationshipRegistry types.ClientRegistrationRelationshipRegistry
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))

	b := store.Get(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRegistryW3CIdentitifer,
	))

	k.cdc.MustUnmarshal(b, &clientRegistrationRelationshipRegistry)

	for index, existingEntry := range clientRegistrationRelationshipRegistry.Relationships {
		if existingEntry.Owner.GetW3CIdentifier() == ownerClientRegistrationW3CIdentitifer && existingEntry.Destination.GetW3CIdentifier() == destinationClientRegistrationW3CIdentitifer {
			clientRegistrationRelationshipRegistry.Relationships = append(clientRegistrationRelationshipRegistry.Relationships[:index], clientRegistrationRelationshipRegistry.Relationships[index+1:]...)
		}
	}

	b = k.cdc.MustMarshal(&clientRegistrationRelationshipRegistry)

	store.Set(types.ClientRegistrationRelationshipRegistryKey(
		clientRegistrationRegistryW3CIdentitifer,
	), b)

	return nil

}
func (k Keeper) GetAllClientRegistrationRelationshipRegistry(ctx sdk.Context) (list []types.ClientRegistrationRelationshipRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRelationshipRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientRegistrationRelationshipRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
