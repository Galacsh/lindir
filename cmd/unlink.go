package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:   constants.CMD_UNLINK,
	Short: unlinkCmdShort(),
	Long:  unlinkCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_UNLINK)
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
