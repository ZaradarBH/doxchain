package cli

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListPartitionedPoolRegistries() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-partitioned-pool-registries",
		Short: "list all PartitionedPoolRegistries",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			pageReq, err := client.ReadPageRequest(cmd.Flags())

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PartitionedPoolRegistryAll(context.Background(), &types.QueryAllPartitionedPoolRegistriesRequest{
				Pagination: pageReq,
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPartitionedPoolRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-partitioned-pool-registry [partitioned-pool-registry-w3c-identifier]",
		Short: "shows a PartitionedPoolRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PartitionedPoolRegistry(context.Background(), &types.QueryGetPartitionedPoolRegistryRequest{
				PartitionedPoolRegistryW3CIdentifier: args[0],
			})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
