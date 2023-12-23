package app

import (
	"lindir/app/check"
	"lindir/app/connector"
	"lindir/common/types"
)

// Hard link 'from' directory to 'to' directory.
// Hard linking directories means just connecting both directories.
func (l lindir) Link(from, to types.Path) error {
	// 'from' must be initialized
	err := check.ErrIfNotInitialized(from)
	if err != nil {
		return err
	}

	// 'to' must not be initialized
	notInitialized, err := check.IsNotInitialized(to)
	if err != nil {
		return err
	}

	// initialize connector based on 'from'
	connector, err := connector.NewConnector(from)
	if err != nil {
		return err
	}

	if notInitialized {
		// link 'from' to 'to'
		err = linkAppDir(from, to)
		if err != nil {
			return err
		}

		err = linkTrackedFiles(from, to)
		if err != nil {
			return err
		}

		connector.Connect(to)
	} else {
		err = connector.ErrIfConnected(to)
		if err != nil {
			return err
		}
	}

	return nil
}

func linkAppDir(from, to types.Path) error {
	return &notImplementedError{}
}

func linkTrackedFiles(from, to types.Path) error {
	return &notImplementedError{}
}

func linkFile(relPath string, from, to types.Path) error {
	return &notImplementedError{}
}
