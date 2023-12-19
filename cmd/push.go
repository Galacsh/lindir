package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   constants.CMD_PUSH,
	Short: pushCmdShort(),
	Long:  pushCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_PUSH)
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
