package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInterchainAccount invokes InitInterchainAccount which binds a new port for the account owner and initiates the ics27 channel handshake
func (k Keeper) RegisterInterchainAccount(
	ctx sdk.Context,
	owner sdk.AccAddress,
	connectionID string,
	counterpartyConnectionID string,
) error {
	if err := k.icaKeeper.InitInterchainAccount(ctx, connectionID, counterpartyConnectionID, owner.String()); err != nil {
		return err
	}

	return nil
}
