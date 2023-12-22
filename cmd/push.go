package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   constants.CMD_PUSH,
	Short: pushCmdShort(),
	Long:  pushCmdLong(),
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var targetDir string
		var err error

		if len(args) == 0 {
			targetDir, err = os.Getwd()
		} else {
			targetDir, err = filepath.Abs(args[0])
		}

		if err != nil {
			return &cannotGetDirectory{constants.CMD_PUSH, err}
		}

		err = lindir.Push(types.Path(targetDir))
		if err != nil {
			return &pushError{targetDir, err}
		}

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
