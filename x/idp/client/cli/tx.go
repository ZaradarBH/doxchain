package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/be-heroes/doxchain/x/idp/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAuthenticationRequest())
	cmd.AddCommand(CmdCreateClientRegistrationRegistry())
	cmd.AddCommand(CmdUpdateClientRegistrationRegistry())
	cmd.AddCommand(CmdDeleteClientRegistrationRegistry())
	cmd.AddCommand(CmdCreateClientRegistration())
	cmd.AddCommand(CmdUpdateClientRegistration())
	cmd.AddCommand(CmdDeleteClientRegistration())
	cmd.AddCommand(CmdCreateClientRegistrationRelationship())
	cmd.AddCommand(CmdDeleteClientRegistrationRelationship())	

	return cmd
}
