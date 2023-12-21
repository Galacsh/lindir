package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   constants.CMD_LINK + " [<from>] <to>",
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
	return "Link your working directory to another directory"
}

func linkCmdLong() string {
	description := `
This command will create a connection from one directory to another.

Arguments:
* <from> (optional): The directory to link from. If not provided, the current
	working directory will be used.
* <to>: The directory to link to.
`
	return strings.TrimSpace(description)
}
