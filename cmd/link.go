package cmd

import (
	"lindir/common/constants"
	"lindir/common/types"
	"strings"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   constants.CMD_LINK + " [<from>] <to>",
	Short: linkCmdShort(),
	Long:  linkCmdLong(),
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var from, to types.Path
		var ferr, terr error

		// assign fromDir and toDir
		if len(args) == 1 {
			from, ferr = types.Path(".").Abs()
			to, terr = types.Path(args[0]).Abs()
		} else {
			from, ferr = types.Path(args[0]).Abs()
			to, terr = types.Path(args[1]).Abs()
		}

		if ferr != nil {
			return &cannotGetDirectory{constants.CMD_LINK, ferr}
		} else if terr != nil {
			return &cannotGetDirectory{constants.CMD_LINK, terr}
		}

		err := lindir.Link(from, to)
		if err != nil {
			return &linkError{from.String(), to.String(), err}
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
