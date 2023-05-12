package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/dgrijalva/jwt-go"
)

// Login method for simple idp keeper
func (k Keeper) Login(ctx sdk.Context, msg types.MsgAuthenticationRequest) (types.MsgAuthenticationResponse, error) {
	response := types.MsgAuthenticationResponse{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.IdpMasterKey)
	idpMasterKeyBytes := store.Get(byteKey)

	if idpMasterKeyBytes == nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not authenticate user")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": msg.Creator,
		"exp": ctx.BlockTime().Add(time.Hour * 1),
	})

	tokenString, err := jwtToken.SignedString([]byte(idpMasterKeyBytes))
	if err != nil {
		return response, sdkerrors.Wrap(types.LoginError, "Could not authenticate user")
	}

	response.Token = tokenString

	return response, nil
}
