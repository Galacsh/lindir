package app

import (
	"lindir/app/setup"
	"lindir/common/types"
)

func (l lindir) Init(dir types.Path) error {
	err := setup.ErrIfAlreadyInitialized(dir)
	if err != nil {
		return err
	}

	err = setup.CreateAppDir(dir)
	if err != nil {
		return err
	}

	err = setup.CreateTrackerFile(dir)
	if err != nil {
		return err
	}

	return setup.CreateConnectorFile(dir)
}
