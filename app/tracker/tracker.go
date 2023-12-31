package tracker

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type tracker struct {
	base     types.Path
	file     types.Path
	tracking types.PathSet
}

// Returns the tracker file of the given directory
func TrackerFileOf(dir types.Path) types.Path {
	return dir.Join(constants.TRACKER)
}

// Creates a new tracker file
func CreateTrackerFile(dir types.Path) error {
	trackerFile := TrackerFileOf(dir)
	return trackerFile.Create()
}

// Returns a new tracker.
// The tracker contains the paths of the tracking files.
func NewTracker(dir types.Path) (*tracker, error) {
	file := TrackerFileOf(dir)
	trackingFiles, err := file.Read()
	if err != nil {
		return nil, err
	}

	return &tracker{dir, file, trackingFiles}, nil
}

// Returns the tracking files
func (t tracker) TrackingFiles() types.PathSet {
	return t.tracking
}

// Returns true if the given file is being tracked
func (t tracker) IsTracking(file string) bool {
	return t.tracking.Contains(file)
}

// Track the given file.
// This function does not save the tracker file.
func (t *tracker) Track(file string) {
	t.tracking.Add(file)
}

// Untrack the given file.
// This function does not save the tracker file.
func (t *tracker) Untrack(file string) {
	t.tracking.Remove(file)
}

// Save tracking files to the tracker file
func (t tracker) Save() error {
	return t.file.Write(t.tracking)
}
