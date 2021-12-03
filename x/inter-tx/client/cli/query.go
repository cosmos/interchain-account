package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the inter-tx module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(getIBCAccountCmd())

	return cmd
}

// getIBCAccountCmd builds a cobra command to query for an interchain account registered on this chain
func getIBCAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "interchainaccounts [account] [connectionId] [counterpartyConnectionId]",
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner := args[0]
			connectionId := args[1]
			counterpartyConnectionId := args[2]

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.InterchainAccountFromAddress(
				context.Background(),
				&types.QueryInterchainAccountFromAddressRequest{Owner: owner, ConnectionId: connectionId, CounterpartyConnectionId: counterpartyConnectionId},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
