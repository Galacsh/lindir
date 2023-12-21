package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   constants.CMD_LINK + " [from] [to]",
	Short: linkCmdShort(),
	Long:  linkCmdLong(),
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var fromDir, toDir string
		var err error

		// assign fromDir and toDir
		if len(args) == 1 {
			fromDir, err = os.Getwd()
			if err != nil {
				return &cannotGetWorkingDir{constants.CMD_LINK, err}
			}
			toDir = args[0]
		} else if len(args) == 2 {
			fromDir = args[0]
			toDir = args[1]
		}

		from, err := types.Path(fromDir).Abs()
		if err != nil {
			return &linkError{fromDir, toDir, err}
		}

		to, err := types.Path(toDir).Abs()
		if err != nil {
			return &linkError{fromDir, toDir, err}
		}

		err = lindir.Link(from, to)
		if err != nil {
			return &linkError{fromDir, toDir, err}
		}

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
