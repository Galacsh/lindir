package cmd

import (
	"fmt"
	"lindir/common/colors"
	"lindir/common/constants"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   constants.CMD_STATUS,
	Short: statusCmdShort(),
	Long:  statusCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetWorkingDir{constants.CMD_STATUS, err}
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
	return ""
}

func statusCmdLong() string {
	return ""
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
