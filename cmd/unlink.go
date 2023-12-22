package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   constants.CMD_UNLINK,
	Short: unlinkCmdShort(),
	Long:  unlinkCmdLong(),
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
			return &cannotGetDirectory{constants.CMD_UNLINK, err}
		}

		err = lindir.Unlink(types.Path(targetDir))
		if err != nil {
			return &unlinkError{targetDir, err}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}

func unlinkCmdShort() string {
	return "Unlink the directory and make it a whole new copy"
}

func unlinkCmdLong() string {
	description := `
This command will unlink the directory and make it a whole new copy.

It will make a copy of the directory, remove the original directory, and rename
the copied directory to the original directory name.

Notice that after making a whole new copy, the directory will not contain 
'{{APP_DIR}}' directory.
`

	description = strings.ReplaceAll(description, "{{APP_DIR}}", constants.APP_DIR)
	return strings.TrimSpace(description)
}

// TODO: make commands with no arguments to support at most 1 argument
