package app

import (
	"io/fs"
	"lindir/common/types"
	"path/filepath"
)

func (l lindir) Status(dir types.Path) (types.PathSet, types.PathSet, error) {
	// working directory must be initialized
	notInitialized, err := isNotInitialized(dir)
	if err != nil {
		return nil, nil, err
	}

	if notInitialized {
		return nil, nil, &notInitializedError{dir}
	}

	// get tracked files
	tracker, err := newTracker(dir)
	if err != nil {
		return nil, nil, err
	}

	// get ignore patterns
	ignorePatterns, err := ignorePatternsOf(dir)
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
		for ignore := range ignorePatterns {
			matched, err := filepath.Match(ignore, relPath)
			if err != nil {
				return err
			}

			if matched {
				if d.IsDir() {
					return fs.SkipDir
				}

				return nil
			}
		}

		// nothing to do if it's directory
		if d.IsDir() {
			return nil
		}

		if tracker.isTracking(relPath) {
			notDeleted.Add(relPath)
		} else {
			added.Add(relPath)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	// remaining files in 'trackedFiles' are deletedFiles
	deleted := tracker.files.Difference(notDeleted)

	return added, deleted, nil
}
