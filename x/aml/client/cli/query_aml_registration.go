package cli

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowAMLRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-aml-request",
		Short: "shows AMLRegistration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetAMLRegistrationRequest{}
			res, err := queryClient.AMLRegistration(context.Background(), params)
			
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
