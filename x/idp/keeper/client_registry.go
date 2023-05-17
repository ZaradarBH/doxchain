package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetClientRegistry set a specific ClientRegistry in the store from its index
func (k Keeper) SetClientRegistry(ctx sdk.Context, ClientRegistry types.ClientRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&ClientRegistry)
	store.Set(types.ClientRegistryKey(
		ClientRegistry.Creator,
	), b)
}

// GetClientRegistry returns a ClientRegistry from its index
func (k Keeper) GetClientRegistry(
	ctx sdk.Context,
	creator string,

) (val types.ClientRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))

	b := store.Get(types.ClientRegistryKey(
		creator,
	))
	
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClientRegistry removes a ClientRegistry from the store
func (k Keeper) RemoveClientRegistry(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	store.Delete(types.ClientRegistryKey(
		creator,
	))
}

// GetAllClientRegistry returns all ClientRegistry
func (k Keeper) GetAllClientRegistry(ctx sdk.Context) (list []types.ClientRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetClientRegistration(ctx sdk.Context, clientRegistration types.ClientRegistration) {
	clientRegistry, found := k.GetClientRegistry(ctx, clientRegistration.Id.Creator)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetFullyQualifiedDidIdentifier() == clientRegistration.Id.GetFullyQualifiedDidIdentifier() {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				break
			}
		}

		clientRegistry.Registrations = append(clientRegistry.Registrations, &clientRegistration)

		k.SetClientRegistry(ctx, clientRegistry)
	}
}

func (k Keeper) RemoveClientRegistration(ctx sdk.Context, creator string, name string) {
	clientRegistry, found := k.GetClientRegistry(ctx, creator)

	if found {
		for index, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Name == name {
				clientRegistry.Registrations = append(clientRegistry.Registrations[:index], clientRegistry.Registrations[index+1:]...)

				k.SetClientRegistry(ctx, clientRegistry)

				break
			}
		}
	}
}

func (k Keeper) GetClientRegistration(ctx sdk.Context, creator string, fullyQualifiedDidIdentifier string) types.ClientRegistration {
	clientRegistry, found := k.GetClientRegistry(ctx, creator)

	if found {
		for _, existingClientRegistration := range clientRegistry.Registrations {
			if existingClientRegistration.Id.GetFullyQualifiedDidIdentifier() == fullyQualifiedDidIdentifier {
				return *existingClientRegistration
			}
		}
	}
	
	return types.ClientRegistration{}
}

func (k Keeper) SetClientRegistrationRelationship(ctx sdk.Context, clientRegistrationRelationship types.ClientRegistrationRelationship) error {
	ownerRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationship.OwnerId.Creator, clientRegistrationRelationship.OwnerId.GetFullyQualifiedDidIdentifier())

	if &ownerRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid owner")
	}

	destinationRegistration := k.GetClientRegistration(ctx, clientRegistrationRelationship.DestinationId.Creator, clientRegistrationRelationship.DestinationId.GetFullyQualifiedDidIdentifier())

	if &destinationRegistration == nil {
		return sdkerrors.Wrap(types.AccessClientListError, "Invalid destination")
	}

	for _, aclEntry := range clientRegistrationRelationship.AccessClientList.Entries {
		matchOwner := false;
		matchDestination := false;

		for _, aclEntryOwner := range ownerRegistration.AccessClientList.Entries {
			if aclEntryOwner.User == aclEntry.User {
				matchOwner = true;
			}
		}

		for _, aclEntryDestination := range destinationRegistration.AccessClientList.Entries {
			if aclEntryDestination.User == aclEntry.User {
				matchDestination = true;
			}
		}

		if !matchOwner || !matchDestination {
			return sdkerrors.Wrap(types.AccessClientListError, "Illigal relationship. Acl must match both owner and destination acls")
		}
	}

	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationGraphKeyPrefix))
	// graph := utils.Graph[didTypes.Did, types.AccessClientList]{}
	// jsonData, err := json.Marshal(graph)
	
	// if err != nil {
	// 	return sdkerrors.Wrap(types.AccessClientListError, "Unable to serialize relationship graph")
	// }

	// b := k.cdc.Marshal(&jsonData)
	// store.Set(types.ClientRegistrationGraphKey(
	//  	ownerRegistration.Id.Creator,
	// ), b)
	//TODO: Implement graph mapping, merge & persistance logic
	return nil
}