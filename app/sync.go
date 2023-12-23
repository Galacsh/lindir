package app

import (
	"lindir/app/check"
	"lindir/app/connector"
	"lindir/common/types"
)

// Push changes in each directory to connected directories
func (l lindir) Sync(dir types.Path) error {
	err := check.ErrIfNotInitialized(dir)
	if err != nil {
		return err
	}

	connector, err := connector.NewConnector(dir)
	if err != nil {
		return err
	}

	// at least one directory must be connected
	err = connector.ErrIfNoConnections()
	if err != nil {
		return err
	}

	// push files to connected directories
	for connection := range connector.Connections() {
		err = l.Push(types.Path(connection))
		if err != nil {
			return err
		}
	}

	return nil
}
