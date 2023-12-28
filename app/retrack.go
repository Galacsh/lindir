package app

import (
	"lindir/app/check"
	"lindir/app/ignorepattern"
	"lindir/app/tracker"
	"lindir/common/types"
)

// Remove the files that are tracked but matches any of the ignore patterns.
func (l lindir) Retrack(dir types.Path) (int, error) {
	// working directory must be initialized
	err := check.ErrIfNotInitialized(dir)
	if err != nil {
		return 0, err
	}

	// initialize tracker
	tracker, err := tracker.NewTracker(dir)
	if err != nil {
		return 0, err
	}

	// update tracker at the last even if an error occurs
	defer tracker.Save()

	ignorePatterns, err := ignorepattern.NewIgnorePatterns(dir)
	if err != nil {
		return 0, err
	}

	ignored := 0
	for tracked := range tracker.TrackingFiles() {
		matched, err := ignorePatterns.Match(tracked)
		if err != nil {
			return ignored, err
		}

		if matched {
			tracker.Untrack(tracked)
			ignored++
		}
	}

	return ignored, nil
}
