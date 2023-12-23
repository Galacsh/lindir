package connector

import (
	"fmt"
	"lindir/common/types"
)

type connectedToOtherDirectoriesError struct {
	dir types.Path
}

func (e connectedToOtherDirectoriesError) Error() string {
	return fmt.Sprintf("'%s' is already linked to other directories", e.dir)
}

type alreadyConnectedError struct {
	from types.Path
	to   types.Path
}

func (e alreadyConnectedError) Error() string {
	return fmt.Sprintf("'%s' is already linked to '%s'", e.from, e.to)
}
