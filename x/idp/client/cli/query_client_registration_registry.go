package cli

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListClientRegistrationRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-client-registries",
		Short: "list all ClientRegistrationRegistry",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			pageReq, err := client.ReadPageRequest(cmd.Flags())

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryAllClientRegistrationRegistryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ClientRegistrationRegistryAll(context.Background(), params)

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

func CmdShowClientRegistrationRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-client-registry [creator]",
		Short: "shows a ClientRegistrationRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ClientRegistrationRegistry(context.Background(), &types.QueryGetClientRegistrationRegistryRequest{
				Creator: args[0],
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
