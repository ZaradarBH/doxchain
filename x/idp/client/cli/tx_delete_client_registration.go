package cli

import (
	"strconv"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDeleteClientRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-client-registration [client-registration-registry-w3c-identifier] [client-registration-w3c-identifier]",
		Short: "Broadcast message DeleteClientRegistration",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteClientRegistration(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
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
