package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   constants.CMD_INIT,
	Short: initCmdShort(),
	Long:  initCmdLong(),
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return &cannotGetWorkingDir{}
		}

		return lindir.Init(types.Path(wd))
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
