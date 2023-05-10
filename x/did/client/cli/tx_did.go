package cli

import (	
    "encoding/json"

	"doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did [didjson]",
		Short: "Create a new did",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			
			var did types.Did
			err = json.Unmarshal([]byte(args[0]), &did)
			
			if err != nil {
				return err
			}
			
			msg := types.NewMsgCreateDid(did)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did [didjson]",
		Short: "Update a did",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
						
			var did types.Did
			err = json.Unmarshal([]byte(args[0]), &did)
			
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDid(did)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-did [fullyQualifiedDidIdentifier]",
		Short: "Delete a did by fullyQualifiedDidIdentifier",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDid(clientCtx.GetFromAddress().String(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}