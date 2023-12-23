package app

import (
	"io/fs"
	"lindir/app/check"
	"lindir/app/ignorepattern"
	"lindir/app/tracker"
	"lindir/common/types"
	"path/filepath"
)

func (l lindir) Status(dir types.Path) (types.PathSet, types.PathSet, error) {
	return status(dir)
}

func status(dir types.Path) (types.PathSet, types.PathSet, error) {
	// working directory must be initialized
	err := check.ErrIfNotInitialized(dir)
	if err != nil {
		return nil, nil, err
	}

	// initialize tracker
	tracker, err := tracker.NewTracker(dir)
	if err != nil {
		return nil, nil, err
	}

	// initialize ignore patterns
	ignorePatterns, err := ignorepattern.NewIgnorePatterns(dir)
	if err != nil {
		return nil, nil, err
	}

	added := make(types.PathSet)
	notDeleted := make(types.PathSet)

	err = filepath.WalkDir(dir.String(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir.String(), path)
		if err != nil {
			return err
		}

		// ignore if it matches any ignore pattern
		matched, err := ignorePatterns.Match(relPath)
		if err != nil {
			return err
		}

		if matched {
			if d.IsDir() {
				return fs.SkipDir
			}

			return nil
		}

		// nothing to do if it's directory
		if d.IsDir() {
			return nil
		}

		if tracker.IsTracking(relPath) {
			notDeleted.Add(relPath)
		} else {
			added.Add(relPath)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	deleted := tracker.TrackingFiles().Difference(notDeleted)

	return added, deleted, nil
}
