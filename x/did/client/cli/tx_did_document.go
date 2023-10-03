package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did-document [did-document-json]",
		Short: "Create a DidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			didDocument := types.DidDocument{}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			jsonData, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			}

			err = clientCtx.Codec.UnmarshalJSON(jsonData, &didDocument)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDidDocumentRequest(clientCtx.GetFromAddress().String(), didDocument)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did-document [did-document-json]",
		Short: "Update a DidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			jsonData, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			}

			var didDocument types.DidDocument
			err = clientCtx.Codec.UnmarshalJSON(jsonData, &didDocument)

			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDidDocumentRequest(clientCtx.GetFromAddress().String(), didDocument)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-did-document [did-url]",
		Short: "Delete a DidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDidDocumentRequest(clientCtx.GetFromAddress().String(), args[0])

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
