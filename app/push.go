package app

import "lindir/common/types"

func (l lindir) Push(dir types.Path) error {
	// working directory must be initialized
	notInitialized, err := isNotInitialized(dir)
	if err != nil {
		return err
	}

	if notInitialized {
		return &notInitializedError{dir}
	}

	// get connected directories
	connector, err := newConnector(dir)
	if err != nil {
		return err
	}

	// get status of the working directory
	added, deleted, err := status(dir)
	if err != nil {
		return err
	}

	// initialize tracker
	tracker, err := newTracker(dir)
	if err != nil {
		return err
	}

	// update tracker at the last even if an error occurs
	defer tracker.save()

	for connection := range connector.connections() {
		if connection == dir.String() {
			continue
		}

		connectedDir := types.Path(connection)

		// create hard links for new files
		for file := range added {
			err := linkFile(file, dir, connectedDir)
			if err != nil {
				return err
			}
			tracker.track(file)
		}

		// delete files that were deleted in the working directory
		for file := range deleted {
			fileToDelete := connectedDir.Join(file)

			err = fileToDelete.Remove()
			if err != nil {
				return err
			}
			tracker.untrack(file)
		}
	}

	return nil
}
