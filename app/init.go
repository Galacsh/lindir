package app

import (
	"lindir/app/initializer"
	"lindir/common/types"
)

func (l lindir) Init(dir types.Path) error {
	err := initializer.ErrIfAlreadyInitialized(dir)
	if err != nil {
		return err
	}

	err = initializer.CreateAppDir(dir)
	if err != nil {
		return err
	}

	err = initializer.CreateTrackerFile(dir)
	if err != nil {
		return err
	}

	return initializer.CreateConnectorFile(dir)
}
