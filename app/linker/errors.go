package linker

import (
	"fmt"
	"lindir/common/types"
)

type fileWithSameNameExistsError struct {
	file types.Path
}

func (e fileWithSameNameExistsError) Error() string {
	return fmt.Sprintf("'%s' already exists", e.file)
}
