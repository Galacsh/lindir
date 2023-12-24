package cmd

import (
	"fmt"
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   constants.CMD_PUSH + " [<directory>]",
	Short: pushCmdShort(),
	Long:  pushCmdLong(),
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
			return &cannotGetDirectory{constants.CMD_PUSH, err}
		}

		added, deleted, err := lindir.Status(types.Path(targetDir))
		if err != nil {
			return &pushError{targetDir.String(), err}
		}

		err = lindir.Push(types.Path(targetDir), added, deleted)
		if err != nil {
			return &pushError{targetDir.String(), err}
		}

		afterPush(targetDir, added.Len(), deleted.Len())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func pushCmdShort() string {
	return "Push new/deleted files to linked directories"
}

func pushCmdLong() string {
	description := `
This command will push new/deleted files to linked directories.

This is a one-way directional command. Files that are added/deleted in linked 
directories are not considered.

To sync directories(bi-directional), use '{{CMD}} {{CMD_SYNC}}' instead.
`

	description = strings.ReplaceAll(description, "{{CMD}}", constants.CMD)
	description = strings.ReplaceAll(description, "{{CMD_SYNC}}", constants.CMD_SYNC)
	return strings.TrimSpace(description)

}

func afterPush(targetDir types.Path, added, deleted int) {
	if added == 0 && deleted == 0 {
		fmt.Printf("Nothing to push to linked directories.\n")
		return
	}

	fmt.Printf("Added %d files and deleted %d files in linked directories.\n", added, deleted)
}
