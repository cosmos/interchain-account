package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
)

// TrySendCoins builds a banktypes.NewMsgSend and uses the ibc-account module keeper to send the message to another chain
func (keeper Keeper) TrySendCoins(
	ctx sdk.Context,
	owner sdk.AccAddress,
	fromAddr,
	toAddr string,
	amt sdk.Coins,
	connectionId string,
	counterpartyConnectionId string,
) error {
	portId, err := icatypes.GeneratePortID(owner.String(), connectionId, counterpartyConnectionId)
	if err != nil {
		return err
	}
	chanCap := keeper.icaControllerKeeper.GetCapability(ctx, ibchost.PortPath(portId))

	msg := &banktypes.MsgSend{FromAddress: fromAddr, ToAddress: toAddr, Amount: amt}
	data, err := icatypes.SerializeCosmosTx(keeper.cdc, []sdk.Msg{msg})
	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	_, err = keeper.icaControllerKeeper.TrySendTx(ctx, chanCap, portId, packetData)
	return err
}
