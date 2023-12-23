package app

import (
	"lindir/app/check"
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/types"
)

func (l lindir) Init(dir types.Path) error {
	err := check.ErrIfAlreadyInitialized(dir)
	if err != nil {
		return err
	}

	noTracker, err := check.NoTrackerFile(dir)
	if err != nil {
		return err
	}

	if noTracker {
		err = tracker.CreateTrackerFile(dir)
		if err != nil {
			return err
		}
	}

	noConnector, err := check.NoConnectorFile(dir)
	if err != nil {
		return err
	}

	if noConnector {
		err = connector.CreateConnectorFile(dir)
		if err != nil {
			return err
		}
	}

	return nil
}
