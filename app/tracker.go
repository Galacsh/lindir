package app

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type tracker struct {
	base     types.Path
	file     types.Path
	tracking types.PathSet
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

func (t tracker) trackingFiles() types.PathSet {
	return t.tracking
}

// Returns true if the given file is being tracked
func (t tracker) isTracking(file string) bool {
	return t.tracking.Contains(file)
}

func (t *tracker) track(file string) {
	t.tracking.Add(file)
}

func (t *tracker) untrack(file string) {
	t.tracking.Remove(file)
}

func (t tracker) save() error {
	return t.file.Write(t.tracking)
}
