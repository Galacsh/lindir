package app

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type tracker struct {
	base  types.Path
	files types.PathSet
}

// Creates a new tracker file
func createTracker(dir types.Path) error {
	return &notImplementedError{}
}

// Returns a new tracker
func newTracker(dir types.Path) (*tracker, error) {
	trackerFile := dir.Join(constants.TRACKER)
	files, err := trackerFile.Read()
	if err != nil {
		return nil, err
	}

	return &tracker{base: dir, files: files}, nil
}

// Returns true if the given file is being tracked
func (t tracker) isTracking(file string) bool {
	return t.files.Contains(file)
}
