package cmd

import (
	"fmt"
	"lindir/common/constants"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   constants.CMD_LINK,
	Short: linkCmdShort(),
	Long:  linkCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%v called", constants.CMD_LINK)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}

func linkCmdShort() string {
	return ""
}

func linkCmdLong() string {
	return ""
}
