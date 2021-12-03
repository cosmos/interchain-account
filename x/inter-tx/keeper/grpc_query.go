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
func (k Keeper) InterchainAccountFromAddress(ctx context.Context, req *types.QueryInterchainAccountFromAddressRequest) (*types.QueryInterchainAccountFromAddressResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	portID, err := icatypes.GeneratePortID(req.Owner, req.ConnectionId, req.CounterpartyConnectionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(sdkCtx, portID)
	if found == false {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	interchainAcc := types.QueryInterchainAccountFromAddressResponse{InterchainAccountAddress: addr}

	return &interchainAcc, nil
}
