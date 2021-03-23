package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
)

const (
	TypeMsgRegisterAccount = "register"
	TypeMsgSend            = "send"
)

var _ sdk.Msg = &MsgRegisterAccount{}

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterAccount(
	port, channel string, height clienttypes.Height, timestamp uint64, owner string,
) *MsgRegisterAccount {
	return &MsgRegisterAccount{
		SourcePort:       port,
		SourceChannel:    channel,
		TimeoutHeight:    height,
		TimeoutTimestamp: timestamp,
		Owner:            owner,
	}
}

// Route implements sdk.Msg
func (MsgRegisterAccount) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgRegisterAccount) Type() string {
	return TypeMsgRegisterAccount
}

func (msg MsgRegisterAccount) ValidateBasic() error {
	if msg.Owner == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	return nil
}

func (msg MsgRegisterAccount) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterAccount) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var _ sdk.Msg = &MsgSend{}

// NewMsgSend creates a new MsgSend instance
func NewMsgSend(
	chainType, port, channel string, height clienttypes.Height, timestamp uint64, sender, toAddress string, amount []sdk.Coin,
) *MsgSend {
	return &MsgSend{
		ChainType:        chainType,
		SourcePort:       port,
		SourceChannel:    channel,
		TimeoutHeight:    height,
		TimeoutTimestamp: timestamp,
		Sender:           sender,
		ToAddress:        toAddress,
		Amount:           amount,
	}
}

// Route implements sdk.Msg
func (MsgSend) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSend) Type() string {
	return TypeMsgSend
}

// GetSigners implements sdk.Msg
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// ValidateBasic performs a basic check of the MsgRegisterAccount fields.
func (msg MsgSend) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if msg.ToAddress == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

func (msg MsgSend) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}