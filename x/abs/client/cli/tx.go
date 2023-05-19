package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/client"
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

	cmd.AddCommand(CmdUpdateBreakFactor())
	cmd.AddCommand(CmdCreatePartitionedPoolRegistry())

	return cmd
}
