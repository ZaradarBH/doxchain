package cli

import (
	"strconv"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetDotWellKnown() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-dot-well-known [did-url]",
		Short: "Query GetDotWellKnown",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetDotWellKnown(cmd.Context(), &types.QueryGetDotWellKnownRequest{
				TenantW3CIdentifier: args[0],
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
