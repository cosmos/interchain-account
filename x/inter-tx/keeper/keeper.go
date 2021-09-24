package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcacckeeper "github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/keeper"
)

type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	icaKeeper ibcacckeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, iaKeeper ibcacckeeper.Keeper) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		icaKeeper: iaKeeper,
	}
}
