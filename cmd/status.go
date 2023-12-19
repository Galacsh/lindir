package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   constants.CMD_STATUS,
	Short: statusCmdShort(),
	Long:  statusCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_STATUS)
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
