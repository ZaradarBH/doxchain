package cli

import (
	"strconv"

	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDeleteClientRegistrationRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-client-registration-relationship [client-registration-registry-did-url] [owner-client-registration-did-url] [destination-client-registration-did-url]",
		Short: "Broadcast message DeleteClientRegistrationRelationshipRequest",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteClientRegistrationRelationshipRequest(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
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
