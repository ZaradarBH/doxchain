package keeper

import (
	"time"
	"crypto/rand"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/zalando/go-keyring"
)

// Login method for simple idp keeper
func (k Keeper) Login(ctx sdk.Context, msg types.MsgAuthenticationRequest) (types.MsgAuthenticationResponse, error) {
	response := types.MsgAuthenticationResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	//TODO: Need to reconsider this approach. The gpkey will be diff on various validators. If validator X processes block Y when no key exists it would write the new key to the store and on the next block it should be available on to all other nodes and thus the next block producer (theoretically)
	//TODO: Figure out how to protected the key so it cannot be exposed. Can a keeper somehow have a pkey it can use to encrypt data without exposing the pkey?
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IdpMasterKey)
	idpMasterKeyBytes := store.Get(byteKey)

	if idpMasterKeyBytes == nil {
		idpMasterKeyBytes, err := keyring.GeneratePrivateKey("rsa", 2048)
		
		if err != nil {
			return response, sdkerrors.Wrap(types.LoginError, "Failed to initialize new master key")
		}

		store.Set(byteKey, idpMasterKeyBytes)
	}

	//TODO: Move JwtTokenFactory to common util namespace, implement option for passing in claims in factory and replace this logic
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": msg.Creator,
		"exp": ctx.BlockTime().Add(time.Hour * 1),
	})

	tokenString, err := jwtToken.SignedString([]byte(idpMasterKeyBytes))
	if err != nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not issue refresh token")
	}

	response.Token = tokenString

	return response, nil
}

// AuthorizeCreator checks if a creator belongs to a given tenant
func (k Keeper) AuthorizeCreator(ctx sdk.Context, tenant string, creator string) (bool, error) {
	acl, err := k.GetAccessClientList(ctx, tenant)

	if err != nil {
		return false, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.Creator == creator {
			return true, nil
		}
	}

	return false, nil
}
