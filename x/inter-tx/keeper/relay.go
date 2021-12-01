package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
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

	chanId, found := keeper.icaControllerKeeper.GetActiveChannelID(ctx, portId)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portId)
	}

	chanCap, found := keeper.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portId, chanId))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	msg := &banktypes.MsgSend{FromAddress: fromAddr, ToAddress: toAddr, Amount: amt}
	data, err := icatypes.SerializeCosmosTx(keeper.cdc, []sdk.Msg{msg})
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	_, err = keeper.icaControllerKeeper.TrySendTx(ctx, chanCap, portId, packetData)
	return err
}
