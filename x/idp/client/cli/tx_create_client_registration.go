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

func CmdCreateClientRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client-registration [client-registration-json]",
		Short: "Broadcast message CreateClientRegistration",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var clientRegistration types.ClientRegistration
			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &clientRegistration)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClientRegistration(
				clientRegistration,
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
