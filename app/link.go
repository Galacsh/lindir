package app

import (
	"lindir/common/types"
)

// Hard link 'from' directory to 'to' directory.
// Hard linking directories means just connecting both directories.
func (l lindir) Link(from, to types.Path) error {
	// 'from' must be initialized
	fromNotInitialized, err := isNotInitialized(from)
	if err != nil {
		return err
	} else if fromNotInitialized {
		return &notInitializedError{from}
	}

	// 'to' must not be initialized
	toNotInitialized, err := isNotInitialized(to)
	if err != nil {
		return err
	}

	// initialize connector based on 'from'
	connector, err := newConnector(from)
	if err != nil {
		return err
	}

	if toNotInitialized {
		// link 'from' to 'to'
		err = linkAppDir(from, to)
		if err != nil {
			return err
		}

		err = linkTrackedFiles(from, to)
		if err != nil {
			return err
		}

		connector.connect(to)
	} else {
		// do nothing if 'to' is already initialized
		isConnected := connector.hasConnection(to)

		if isConnected {
			return alreadyConnectedError{from, to}
		} else {
			return connectedToOtherDirectoriesError{to}
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
