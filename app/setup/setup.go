package setup

import (
	"lindir/app/connect"
	"lindir/app/track"
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

	return connect.ConnectorFileOf(dir).Write(defaultConnections)
}

// Creates a new tracker file
func CreateTrackerFile(dir types.Path) error {
	return track.TrackerFileOf(dir).Write(types.PathSet{})
}

func IsNotInitialized(dir types.Path) (bool, error) {
	return false, nil
}

func ErrIfNotInitialized(dir types.Path) error {
	return &notInitializedError{dir}
}

func ErrIfAlreadyInitialized(dir types.Path) error {
	return &alreadyInitializedError{dir}
}
