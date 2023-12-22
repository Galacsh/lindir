package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   constants.CMD_SYNC,
	Short: syncCmdShort(),
	Long:  syncCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetDirectory{constants.CMD_SYNC, err}
		}

		err = lindir.Sync(types.Path(wd))
		if err != nil {
			return &syncError{wd, err}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func syncCmdShort() string {
	return "Sync new/deleted files between all linked directories"
}

func syncCmdLong() string {
	description := `
This command will sync new/deleted files between all linked directories.

This is a bi-directional command. Files that are added/deleted in each linked 
directories will also be added/deleted to linked directories.

To push only the changes made in current working directory, 
use '{{CMD}} {{CMD_PUSH}}' instead.
`

	description = strings.ReplaceAll(description, "{{CMD}}", constants.CMD)
	description = strings.ReplaceAll(description, "{{CMD_PUSH}}", constants.CMD_PUSH)
	return strings.TrimSpace(description)
}
