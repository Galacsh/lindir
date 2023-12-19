package cmd

import (
	"lindir/app"
	"lindir/common/types"
	"os"

	"github.com/spf13/cobra"
)

type App interface {
	Init(dir types.Path) error
	Status(dir types.Path) (types.PathSet, types.PathSet, error)
	Link(fromDir, toDir types.Path) error
	Unlink(dir types.Path) error
	Push(dir types.Path) error
	Sync(dir types.Path) error
}

var lindir App

var rootCmd = &cobra.Command{
	Use:   "lindir",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	// initialize lindir
	lindir = app.New()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
