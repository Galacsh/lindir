package cmd

import "fmt"

type cannotGetWorkingDir struct {
	err error
}

func (e cannotGetWorkingDir) Error() string {
	return fmt.Sprintf("cannot get working directory: %s", e.err.Error())
}
