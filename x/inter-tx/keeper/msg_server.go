package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Register implements the Msg/Register gRPC method
func (k msgServer) Register(c context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	acc, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.RegisterInterchainAccount(ctx, acc, msg.ConnectionId, msg.ConnectionId); err != nil {
		return nil, err
	}

	return &types.MsgRegisterAccountResponse{}, nil
}

// Send implements the Msg/Send gRPC method
// Send is used to send tokens from an interchain account to another account on a target chain
// The inter-tx module keeper uses the interchain accounts module keeper to build and send an IBC packet of type EXECUTE_TX
func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.TrySendCoins(ctx, msg.Owner, msg.InterchainAccount, msg.ToAddress, msg.Amount, msg.ConnectionId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "send coins failed: %v", err)
	}

	return &types.MsgSendResponse{}, nil
}
