package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   constants.CMD_INIT,
	Short: initCmdShort(),
	Long:  initCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetDirectory{constants.CMD_INIT, err}
		}

		err = lindir.Init(types.Path(wd))
		if err != nil {
			return &initError{wd, err}
		}

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
