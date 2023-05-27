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

const (
	FlagGrantType = "grant-type"
)

func CmdToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [tenant] [client-id] [client-secret] [scope] [device-code] [authorization-code] [client-assertion] [client-assertion-type]",
		Short: "Broadcast message token",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			grantType, _ := cmd.Flags().GetInt32(FlagGrantType)

			msg := types.NewMsgTokenRequest(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				[]string{args[3]},
				types.GrantType(grantType),
				args[4],
				args[5],
				args[6],
				args[7],
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32(FlagGrantType, 0, "define grant type (default 0 for undefined), 1: client credentials grant, 2: device code grant, 3: authorization code grant")
	cmd.MarkFlagRequired(FlagGrantType)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
