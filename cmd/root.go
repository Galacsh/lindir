package cmd

import (
	"fmt"
	"lindir/app"
	"lindir/common/colors"
	"lindir/common/constants"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           constants.CMD,
	Short:         rootCmdShort(),
	Long:          rootCmdLong(),
	SilenceErrors: true,
}

func Execute() {
	// initialize lindir
	lindir = app.New()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, colors.Red(err.Error()))
		os.Exit(1)
	}
}

func rootCmdShort() string {
	return "Tool for managing hard links across multiple directories"
}

func rootCmdLong() string {
	description := `
{{APP}} is a tool for managing hard links across multiple directories.

It is designed to address the limitations of traditional hard links and
symbolic links. While hard links allow multiple access points to a single file,
they cannot be used for directories. Symbolic links, on the other hand, can
link directories but may not always be compatible with certain programs.

To overcome these limitations, {{APP}} allows you to mimic directory linking
through a Git-inspired approach. {{APP}} tracks and manages hard links by
creating a '{{APP_DIR}}' directory within your chosen directory. This approach
is particularly beneficial in scenarios where you require a single source of
truth for your files, such as managing dot files in a single Git repository.

Run '{{CMD}} help <command>' to learn more about a specific command.
`

	description = strings.ReplaceAll(description, "{{APP}}", constants.APP)
	description = strings.ReplaceAll(description, "{{CMD}}", constants.CMD)
	description = strings.ReplaceAll(description, "{{APP_DIR}}", constants.APP_DIR)
	return strings.TrimSpace(description)
}
