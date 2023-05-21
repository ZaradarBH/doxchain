package cli

import (
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateClientRegistrationRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client-registry [client-registry-json]",
		Short: "Create a new ClientRegistrationRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var clientRegistry types.ClientRegistrationRegistry
			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &clientRegistry)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClientRegistrationRegistry(clientCtx.GetFromAddress().String(), clientRegistry)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateClientRegistrationRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-client-registry [client-registry-json]",
		Short: "Update a ClientRegistrationRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var clientRegistry types.ClientRegistrationRegistry
			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &clientRegistry)

			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClientRegistrationRegistry(clientCtx.GetFromAddress().String(), clientRegistry)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteClientRegistrationRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-client-registry",
		Short: "Delete a ClientRegistrationRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteClientRegistrationRegistry(clientCtx.GetFromAddress().String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
