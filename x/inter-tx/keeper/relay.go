package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TrySendCoins builds a banktypes.NewMsgSend and uses the ibc-account module keeper to send the message to another chain
func (keeper Keeper) TrySendCoins(ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	typ string,
	fromAddr sdk.AccAddress,
	toAddr sdk.AccAddress,
	amt sdk.Coins,
) error {
	ibcAccount, err := keeper.GetIBCAccount(ctx, fromAddr)
	if err != nil {
		return err
	}

	//TODO: I think we do not need to pass an accAddr to NewMsgSend
	acc, _ := sdk.AccAddressFromBech32(ibcAccount)
	msg := banktypes.NewMsgSend(acc, toAddr, amt)

	_, err = keeper.iaKeeper.TryRunTx(ctx, fromAddr, msg)
	return err
}
