package cli

import (
    "context"
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/be-heroes/doxchain/x/abs/types"
)

func CmdListPartitionedPoolRegistries() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-partitioned-pool-registries",
		Short: "list all partitionedPoolRegistries",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllPartitionedPoolRegistriesRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.PartitionedPoolRegistryAll(context.Background(), params)
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
		Use:   "show-partitioned-pools [creator]",
		Short: "shows a partitionedPools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argIndex := args[0]
            
            params := &types.QueryGetPartitionedPoolRegistryRequest{
                Creator: argIndex,
                
            }

            res, err := queryClient.PartitionedPoolRegistry(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
