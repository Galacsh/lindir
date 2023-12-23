package check

import (
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/types"
)

func NoTrackerFile(dir types.Path) (bool, error) {
	return tracker.TrackerFileOf(dir).NotExists()
}

func NoConnectorFile(dir types.Path) (bool, error) {
	return connector.ConnectorFileOf(dir).NotExists()
}

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

func isInitialized(dir types.Path) (bool, error) {
	notInitialized, err := IsNotInitialized(dir)
	if err != nil {
		return false, err
	}

	return !notInitialized, nil
}

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

func ErrIfAlreadyInitialized(dir types.Path) error {
	initialized, err := isInitialized(dir)
	if err != nil {
		return err
	}

	if !initialized {
		return &alreadyInitializedError{dir}
	} else {
		return nil
	}
}
