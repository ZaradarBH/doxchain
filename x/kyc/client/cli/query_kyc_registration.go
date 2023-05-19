package cli

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowKYCRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-kyc-request",
		Short: "shows KYCRegistration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetKYCRegistrationRequest{}

			res, err := queryClient.KYCRegistration(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
