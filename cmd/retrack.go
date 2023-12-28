package cmd

import (
	"fmt"
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var retrackCmd = &cobra.Command{
	Use:   constants.CMD_RETRACK + " [<directory>]",
	Short: retrackCmdShort(),
	Long:  retrackCmdLong(),
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
			return &cannotGetDirectory{constants.CMD_RETRACK, err}
		}

		ignored, err := lindir.Retrack(types.Path(targetDir))
		if err != nil {
			return &retrackError{targetDir.String(), err}
		}

		afterRetrack(ignored)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(retrackCmd)
}

func retrackCmdShort() string {
	return "Removes files from tracking if they match any ignore pattern."
}

func retrackCmdLong() string {
	description := `
This command will remove files from tracking if they match any ignore pattern.

This is useful when you want updated ignore patterns to ignore the files that 
are already tracked.
`

	description = strings.ReplaceAll(description, "{{APP}}", constants.APP)
	description = strings.ReplaceAll(description, "{{APP_DIR}}", constants.APP_DIR)
	return strings.TrimSpace(description)
}

func afterRetrack(ignored int) {
	msg := strings.ReplaceAll("Ignored %d files from '{{TRACKER}}'.\n", "{{TRACKER}}", constants.TRACKER)
	fmt.Printf(msg, ignored)
}
