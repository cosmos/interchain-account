package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/controller/keeper"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	icaControllerKeeper icacontrollerkeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper icacontrollerkeeper.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		icaControllerKeeper: iaKeeper,
	}
}
