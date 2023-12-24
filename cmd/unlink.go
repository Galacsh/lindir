package cmd

import (
	"fmt"
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   constants.CMD_UNLINK + " [<directory>]",
	Short: unlinkCmdShort(),
	Long:  unlinkCmdLong(),
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
			return &cannotGetDirectory{constants.CMD_UNLINK, err}
		}

		err = lindir.Unlink(types.Path(targetDir))
		if err != nil {
			return &unlinkError{targetDir.String(), err}
		}

		afterUnlink(targetDir)

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

func afterUnlink(targetDir types.Path) {
	fmt.Printf("Unlinked '%s'.\n", targetDir.String())
}
