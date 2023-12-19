package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   constants.CMD_SYNC,
	Short: syncCmdShort(),
	Long:  syncCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_SYNC)
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
