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
		Use:   "create-partitioned-pool-registry [creator]",
		Short: "Create a new partitionedPools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         	creator := args[0]
        
            // Get value arguments
		
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePartitionedPoolRegistryRequest(
			    clientCtx.GetFromAddress().String(),
			    creator,
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
