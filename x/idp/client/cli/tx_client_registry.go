package cli

import (
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateClientRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client-registry [client-registry-json]",
		Short: "Create a new ClientRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var clientRegistry types.ClientRegistry
			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &clientRegistry)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClientRegistry(clientRegistry)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateClientRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-client-registry [client-registry-json]",
		Short: "Update a ClientRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var clientRegistry types.ClientRegistry
			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &clientRegistry)

			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClientRegistry(clientRegistry)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteClientRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-client-registry",
		Short: "Delete a ClientRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteClientRegistry(clientCtx.GetFromAddress().String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
