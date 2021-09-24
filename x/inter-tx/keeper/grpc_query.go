package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// IBCAccountFromAddress implements the Query/IBCAccountFromAddress gRPC method
func (k Keeper) IBCAccountFromAddress(c context.Context, req *types.QueryIBCAccountFromAddressRequest) (*types.QueryIBCAccountFromAddressResponse, error) {

	portID, err := icatypes.GeneratePortID(req.Address.String(), req.ConnectionId, "")
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to generate port identifier: %v", err)
	}

	ctx := sdk.UnwrapSDKContext(c)
	addr, found := k.icaKeeper.GetInterchainAccountAddress(ctx, portID)
	if !found {
		return nil, status.Error(codes.NotFound, "failed to retrieve interchain account address")
	}

	return types.NewQueryIBCAccountFromAddressResponse(addr), nil
}
