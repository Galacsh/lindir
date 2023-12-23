package initializer

import (
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/constants"
	"lindir/common/types"
)

func appDirOf(dir types.Path) types.Path {
	return dir.Join(constants.APP_DIR)
}

func CreateAppDir(dir types.Path) error {
	return appDirOf(dir).CreateDir()
}

func CreateConnectorFile(dir types.Path) error {
	defaultConnections := types.PathSet{}
	defaultConnections.AddPath(dir)

	return connector.ConnectorFileOf(dir).Write(defaultConnections)
}

// Creates a new tracker file
func CreateTrackerFile(dir types.Path) error {
	return tracker.TrackerFileOf(dir).Write(types.PathSet{})
}

func noTrackerFile(dir types.Path) (bool, error) {
	return tracker.TrackerFileOf(dir).NotExists()
}

func noConnectorFile(dir types.Path) (bool, error) {
	return connector.ConnectorFileOf(dir).NotExists()
}

func IsNotInitialized(dir types.Path) (bool, error) {
	noTracker, err := noTrackerFile(dir)
	if err != nil {
		return false, err
	}

	noConnector, err := noConnectorFile(dir)
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
