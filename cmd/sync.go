package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   constants.CMD_SYNC,
	Short: syncCmdShort(),
	Long:  syncCmdLong(),
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetWorkingDir{constants.CMD_SYNC, err}
		}

		err = lindir.Sync(types.Path(wd))
		if err != nil {
			return &syncError{wd, err}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func syncCmdShort() string {
	return ""
}

func syncCmdLong() string {
	return ""
}
