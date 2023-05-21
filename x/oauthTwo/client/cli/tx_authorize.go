package cli

import (
	"strconv"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAuthorize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorize [client-id] [scope]",
		Short: "Broadcast message Authorize",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			
			if err != nil {
				return err
			}

			msg := types.NewMsgAuthorizeRequest(
				clientCtx.GetFromAddress().String(),
				args[0],
				[]string{args[1]},
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
