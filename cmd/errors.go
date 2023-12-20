package cmd

import (
	"fmt"
)

type cannotGetWorkingDir struct {
	cmd string
	err error
}

func (e cannotGetWorkingDir) Error() string {
	return fmt.Sprintf("failed to get working directory for '%s': %s", e.cmd, e.err.Error())
}

type initError struct {
	dir string
	err error
}

func (e initError) Error() string {
	return fmt.Sprintf("failed to initialize '%s': %s", e.dir, e.err.Error())
}
