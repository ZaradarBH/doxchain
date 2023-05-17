package cli

import (
	"strconv"

	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateClientRegistrationRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client-registration-relationship [owner-did] [destination-did] [access-client-list]",
		Short: "Broadcast message CreateClientRegistrationRelationshipRequest",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			var ownerDid didTypes.Did
			var destinationDid didTypes.Did
			var acl idpTypes.AccessClientList

			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &ownerDid)

			if err != nil {
				return err
			}

			err = clientCtx.Codec.UnmarshalJSON([]byte(args[1]), &destinationDid)

			if err != nil {
				return err
			}

			err = clientCtx.Codec.UnmarshalJSON([]byte(args[2]), &acl)

			if err != nil {
				return err
			}

			msg := idpTypes.NewMsgCreateClientRegistrationRelationshipRequest(
				ownerDid,
				destinationDid,
				acl,
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
