package cli

import (
	utils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateAMLRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-aml-request [did-url]",
		Short: "Create AMLRegistration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			did := utils.NewDidTokenFactory().Create(clientCtx.GetFromAddress().String(), args[0])
			msg := types.NewMsgCreateAMLRegistration(clientCtx.GetFromAddress().String(), *did)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAMLRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-aml-request",
		Short: "Delete AMLRegistration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAMLRegistration(clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
