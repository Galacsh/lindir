package check

import (
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/types"
)

// Check if there is no tracker file in the directory
func NoTrackerFile(dir types.Path) (bool, error) {
	return tracker.TrackerFileOf(dir).NotExists()
}

// Check if there is no connector file in the directory
func NoConnectorFile(dir types.Path) (bool, error) {
	return connector.ConnectorFileOf(dir).NotExists()
}

// Check if the directory is not initialized.
// A directory is not initialized if there is no tracker file or connector file.
func IsNotInitialized(dir types.Path) (bool, error) {
	noTracker, err := NoTrackerFile(dir)
	if err != nil {
		return false, err
	}

	noConnector, err := NoConnectorFile(dir)
	if err != nil {
		return false, err
	}

	return noTracker || noConnector, nil
}

// Check if the directory is initialized.
// A directory is initialized if there is both tracker file and connector file.
func isInitialized(dir types.Path) (bool, error) {
	notInitialized, err := IsNotInitialized(dir)
	if err != nil {
		return false, err
	}

	return !notInitialized, nil
}

// Return error if the directory is not initialized.
// A directory is not initialized if there is no tracker file or connector file.
func ErrIfNotInitialized(dir types.Path) error {
	notInitialized, err := IsNotInitialized(dir)
	if err != nil {
		return err
	}

	if notInitialized {
		return &notInitializedError{dir}
	} else {
		return nil
	}
}

// Return error if the directory is initialized.
// A directory is initialized if there is both tracker file and connector file.
func ErrIfAlreadyInitialized(dir types.Path) error {
	initialized, err := isInitialized(dir)
	if err != nil {
		return err
	}

	if initialized {
		return &alreadyInitializedError{dir}
	} else {
		return nil
	}
}
