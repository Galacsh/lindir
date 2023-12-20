package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   constants.CMD_PUSH,
	Short: pushCmdShort(),
	Long:  pushCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetWorkingDir{constants.CMD_PUSH, err}
		}

		err = lindir.Push(types.Path(wd))
		if err != nil {
			return &pushError{wd, err}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func pushCmdShort() string {
	return ""
}

func pushCmdLong() string {
	return ""
}
