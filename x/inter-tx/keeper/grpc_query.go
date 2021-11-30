package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v2/modules/apps/27-interchain-accounts/types"

	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// IBCAccountFromAddress implements the Query/IBCAccount gRPC method
func (k Keeper) IBCAccountFromAddress(ctx context.Context, req *types.QueryIBCAccountFromAddressRequest) (*types.QueryIBCAccountFromAddressResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	portId, err := icatypes.GeneratePortID(req.Address.String(), req.ConnectionId, req.CounterpartyConnectionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	addr, exists := k.icaControllerKeeper.GetInterchainAccountAddress(sdkCtx, portId)
	if exists == false {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portId)
	}

	ibcAccount := types.QueryIBCAccountFromAddressResponse{Address: addr}

	return &ibcAccount, nil
}
