package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgDelegate{}
	_ sdk.Msg = &MsgRegisterAccount{}
	_ sdk.Msg = &MsgSend{}
)

// NewMsgDelegate creates a new MsgDelegate instance
func NewMsgDelegate(owner sdk.AccAddress, amt sdk.Coin, interchainAccAddr, validatorAddr string) *MsgDelegate {
	return &MsgDelegate{
		InterchainAccount: interchainAccAddr,
		Owner:             owner,
		ValidatorAddress:  validatorAddr,
		Amount:            amt,
	}
}

// GetSigners implements sdk.Msg
func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDelegate) ValidateBasic() error {
	if strings.TrimSpace(msg.InterchainAccount) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if strings.TrimSpace(msg.ValidatorAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bech32 validator address: %s", msg.ValidatorAddress)
	}

	return nil
}

// NewMsgRegisterAccount creates a new MsgRegisterAccount instance
func NewMsgRegisterAccount(owner, connectionID string) *MsgRegisterAccount {
	return &MsgRegisterAccount{
		Owner:        owner,
		ConnectionId: connectionID,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterAccount) ValidateBasic() error {
	if strings.TrimSpace(msg.Owner) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterAccount) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

// NewMsgSend creates a new MsgSend instance
func NewMsgSend(owner sdk.AccAddress, amt sdk.Coins, interchainAccAddr, toAddr string) *MsgSend {
	return &MsgSend{
		InterchainAccount: interchainAccAddr,
		Owner:             owner,
		ToAddress:         toAddr,
		Amount:            amt,
	}
}

// GetSigners implements sdk.Msg
func (msg MsgSend) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// ValidateBasic implements sdk.Msg
func (msg MsgSend) ValidateBasic() error {
	if strings.TrimSpace(msg.InterchainAccount) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if strings.TrimSpace(msg.ToAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}
