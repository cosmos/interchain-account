package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/interchain-accounts/x/ibc-account/types"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// RegisterIBCAccount uses the ibc-account module keeper to register an account on a target chain
// An address registration queue is used to keep track of registration requests
func (keeper Keeper) RegisterIBCAccount(
	ctx sdk.Context,
	sender sdk.AccAddress,
) error {
	err := keeper.iaKeeper.InitInterchainAccount(ctx, sender.String())
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent("register-interchain-account"))

	return nil
}

// GetIBCAccount returns an interchain account address
func (keeper Keeper) GetIBCAccount(ctx sdk.Context, owner sdk.AccAddress) (string, error) {
	portId := icatypes.IcaPrefix + strings.TrimSpace(owner.String())
	address, err := keeper.iaKeeper.GetInterchainAccountAddress(ctx, portId)

	return address, err
}

// GetIncrementalSalt increments the Salt value by 1 and returns the Salt
func (keeper Keeper) GetIncrementalSalt(ctx sdk.Context) string {
	kvStore := ctx.KVStore(keeper.storeKey)

	key := []byte("salt")

	salt := types.Salt{
		Salt: 0,
	}

	if kvStore.Has(key) {
		keeper.cdc.MustUnmarshalBinaryBare(kvStore.Get(key), &salt)
		salt.Salt++
	}

	bz := keeper.cdc.MustMarshalBinaryBare(&salt)
	kvStore.Set(key, bz)

	return string(bz)
}
