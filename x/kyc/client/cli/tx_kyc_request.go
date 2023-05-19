package cli

import (
	utils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateKYCRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-kyc-request [did-url]",
		Short: "Create KYCRegistration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			did := utils.NewDidTokenFactory().Create(creator, args[0])
			msg := types.NewMsgCreateKYCRegistration(creator, *did)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteKYCRegistration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-kyc-request",
		Short: "Delete KYCRegistration",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteKYCRegistration(clientCtx.GetFromAddress().String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
