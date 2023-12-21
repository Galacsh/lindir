package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   constants.CMD_UNLINK,
	Short: unlinkCmdShort(),
	Long:  unlinkCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetWorkingDir{constants.CMD_UNLINK, err}
		}

		err = lindir.Unlink(types.Path(wd))
		if err != nil {
			return &unlinkError{wd, err}
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
