package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"vesting/x/dex/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group dex queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListVesting())
	cmd.AddCommand(CmdShowVesting())
	cmd.AddCommand(CmdListSentVesting())
	cmd.AddCommand(CmdShowSentVesting())
	cmd.AddCommand(CmdListTimedoutVesting())
	cmd.AddCommand(CmdShowTimedoutVesting())
	// this line is used by starport scaffolding # 1

	return cmd
}
