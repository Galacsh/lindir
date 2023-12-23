package app

import (
	"fmt"
	"lindir/common/types"
)

type alreadyInitializedError struct {
	dir types.Path
}

func (e alreadyInitializedError) Error() string {
	return fmt.Sprintf("'%s' is already initialized", e.dir)
}

type notInitializedError struct {
	dir types.Path
}

func (e notInitializedError) Error() string {
	return fmt.Sprintf("'%s' is not initialized", e.dir)
}

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
