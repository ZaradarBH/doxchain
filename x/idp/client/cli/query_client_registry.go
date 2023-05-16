package cli

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListClientRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-client-registries",
		Short: "list all ClientRegistry",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			pageReq, err := client.ReadPageRequest(cmd.Flags())

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryAllClientRegistryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClientRegistryAll(context.Background(), params)

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

func CmdShowClientRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-client-registry [creator]",
		Short: "shows a ClientRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetClientRegistryRequest{
				Creator: args[0],
			}

			res, err := queryClient.ClientRegistry(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
