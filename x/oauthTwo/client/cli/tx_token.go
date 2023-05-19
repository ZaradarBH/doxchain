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

func CmdToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [tenant] [client-id] [client-secret] [scope] [grant-type] [device-code] [authorization-code] [client-assertion] [client-assertion-type]",
		Short: "Broadcast message token",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTenant := args[0]
			argClientId := args[1]
			argClientSecret := args[2]
			argScope := args[3]
			argGrantType := args[4]
			argDeviceCode := args[5]
			argAuthorizationCode := args[6]
			argClientAssertion := args[7]
			argClientAssertionType := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTokenRequest(
				clientCtx.GetFromAddress().String(),
				argTenant,
				argClientId,
				argClientSecret,
				argScope,
				argGrantType,
				argDeviceCode,
				argAuthorizationCode,
				argClientAssertion,
				argClientAssertionType,
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
