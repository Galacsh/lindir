package check

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
