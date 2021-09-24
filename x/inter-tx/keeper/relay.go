package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/types"
)

// TrySendCoins builds a banktypes.NewMsgSend and uses the interchain accounts module keeper to send the message to another chain
func (k Keeper) TrySendCoins(
	ctx sdk.Context,
	owner sdk.AccAddress,
	fromAddr,
	toAddr string,
	amount sdk.Coins,
	connectionID string,
) error {
	msg := &banktypes.MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      amount,
	}

	portID, err := icatypes.GeneratePortID(owner.String(), connectionID, "")
	if err != nil {
		return err
	}

	_, err = k.icaKeeper.TrySendTx(ctx, portID, msg)

	return err
}
