package cli

import (
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/be-heroes/doxchain/x/abs/types"
)

func CmdCreatePartitionedPoolRegistry() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-partitioned-pool-registry [denom]",
		Short: "Create a new partitionedPoolRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePartitionedPoolRegistryRequest(
			    clientCtx.GetFromAddress().String(),
			    args[0],
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
