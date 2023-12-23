package app

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type tracker struct {
	base          types.Path
	file          types.Path
	trackingFiles types.PathSet
}

// Creates a new tracker file
func createTrackerFile(dir types.Path) error {
	return trackerFileOf(dir).Write(types.PathSet{})
}

// Returns the tracker file of the given directory
func trackerFileOf(dir types.Path) types.Path {
	return dir.Join(constants.TRACKER)
}

// Returns a new tracker
func newTracker(dir types.Path) (*tracker, error) {
	file := trackerFileOf(dir)
	trackingFiles, err := file.Read()
	if err != nil {
		return nil, err
	}

	return &tracker{dir, file, trackingFiles}, nil
}

// Returns true if the given file is being tracked
func (t tracker) isTracking(file string) bool {
	return t.trackingFiles.Contains(file)
}

func (t tracker) difference(files types.PathSet) types.PathSet {
	return t.trackingFiles.Difference(files)
}
