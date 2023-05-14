package cli

import (
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/be-heroes/doxchain/x/abs/types"
)

func CmdCreatePartitionedPools() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-partitioned-pools [index]",
		Short: "Create a new partitionedPools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePartitionedPools(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
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

func CmdUpdatePartitionedPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-partitioned-pools [index]",
		Short: "Update a partitionedPools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePartitionedPools(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
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

func CmdDeletePartitionedPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-partitioned-pools [index]",
		Short: "Delete a partitionedPools",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
             indexIndex := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePartitionedPools(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
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