package cmd

import (
	"fmt"
	"lindir/common/colors"
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   constants.CMD_STATUS,
	Short: statusCmdShort(),
	Long:  statusCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetDirectory{constants.CMD_STATUS, err}
		}

		added, deleted, err := lindir.Status(types.Path(wd))
		if err != nil {
			return &statusError{wd, err}
		}

		printStatus(added, deleted)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func statusCmdShort() string {
	return "Show the status of the current directory"
}

func statusCmdLong() string {
	description := `
This command will show what files are going to be pushed.

Files that are going to be pushed:
  - Files that are not currently hard linked to any other directory
  - Files that have been deleted from the current directory
`

	description = strings.ReplaceAll(description, "{{APP}}", constants.APP)
	description = strings.ReplaceAll(description, "{{APP_DIR}}", constants.APP_DIR)
	return strings.TrimSpace(description)
}

func printStatus(added types.PathSet, deleted types.PathSet) {
	if len(added) == 0 && len(deleted) == 0 {
		fmt.Println("Nothing to push, working directory clean")
		return
	}

	fmt.Printf("New files(%v):\n", len(added))
	for file := range added {
		fmt.Println("\t" + colors.Green(file))
	}

	fmt.Println()

	fmt.Printf("Deleted files(%v):\n", len(deleted))
	for file := range deleted {
		fmt.Println("\t" + colors.Red(file))
	}
}
