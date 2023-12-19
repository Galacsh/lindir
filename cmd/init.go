package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   constants.CMD_INIT,
	Short: initCmdShort(),
	Long:  initCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_INIT)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initCmdShort() string {
	return ""
}

func initCmdLong() string {
	return ""
}
