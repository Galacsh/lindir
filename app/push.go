package app

import (
	"lindir/app/connector"
	"lindir/app/initializer"
	"lindir/app/tracker"
	"lindir/common/types"
)

func (l lindir) Push(dir types.Path) error {
	// working directory must be initialized
	err := initializer.ErrIfNotInitialized(dir)
	if err != nil {
		return err
	}

	// get connected directories
	connector, err := connector.NewConnector(dir)
	if err != nil {
		return err
	}

	// get status of the working directory
	added, deleted, err := status(dir)
	if err != nil {
		return err
	}

	// initialize tracker
	tracker, err := tracker.NewTracker(dir)
	if err != nil {
		return err
	}

	// update tracker at the last even if an error occurs
	defer tracker.Save()

	for connection := range connector.Connections() {
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
			tracker.Track(file)
		}

		// delete files that were deleted in the working directory
		for file := range deleted {
			fileToDelete := connectedDir.Join(file)

			err = fileToDelete.Remove()
			if err != nil {
				return err
			}
			tracker.Untrack(file)
		}
	}

	return nil
}
