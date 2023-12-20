package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"

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
	return ""
}

func unlinkCmdLong() string {
	return ""
}
