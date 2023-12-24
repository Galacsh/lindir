package cmd

import (
	"fmt"
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   constants.CMD_SYNC + " [<directory>]",
	Short: syncCmdShort(),
	Long:  syncCmdLong(),
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var targetDir types.Path
		var err error

		if len(args) == 0 {
			targetDir, err = types.Path(".").Abs()
		} else {
			targetDir, err = types.Path(args[0]).Abs()
		}

		if err != nil {
			return &cannotGetDirectory{constants.CMD_SYNC, err}
		}

		err = lindir.Sync(types.Path(targetDir))
		if err != nil {
			return &syncError{targetDir.String(), err}
		}

		afterSync(targetDir)

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

func afterSync(targetDir types.Path) {
	fmt.Println("Successfully synchronized.")
}
