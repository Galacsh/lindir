package cmd

import (
	"fmt"
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   constants.CMD_INIT + " [<directory>]",
	Short: initCmdShort(),
	Long:  initCmdLong(),
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
			return &cannotGetDirectory{constants.CMD_INIT, err}
		}

		err = lindir.Init(targetDir)
		if err != nil {
			return &initError{targetDir.String(), err}
		}

		afterInit()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initCmdShort() string {
	return "Initialize a directory for use with " + constants.APP
}

func initCmdLong() string {
	description := `
{{APP}} requires a '{{APP_DIR}}' directory to track and manage hard links.

This command creates '{{APP_DIR}}' directory inside current working directory.
Inside the '{{APP_DIR}}' directory, you will find two files:

* 'connector'
  - This file shows directories that are linked to the current directory.
  - Will be automatically updated when you run the 'link' or 'unlink' commands.
  - Right after initialization, this file will contain current directory.
  - Do not modify this file manually.
* 'tracker'
  - This file shows what files are currently hard linked across directories.
  - Will be automatically updated when you run the 'push' or 'sync' commands.
  - Right after initialization, this file will be empty.
  - Do not modify this file manually.
`
	description = strings.ReplaceAll(description, "{{APP}}", constants.APP)
	description = strings.ReplaceAll(description, "{{APP_DIR}}", constants.APP_DIR)
	return strings.TrimSpace(description)
}

func afterInit() {
	fmt.Println("Successfully initialized.")
}
