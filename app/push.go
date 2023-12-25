package app

import (
	"lindir/app/check"
	"lindir/app/connector"
	"lindir/app/linker"
	"lindir/app/tracker"
	"lindir/common/types"
)

// Push changes in the working directory to connected directories
func (l lindir) Push(dir types.Path, added, deleted types.PathSet) error {
	// working directory must be initialized
	err := check.ErrIfNotInitialized(dir)
	if err != nil {
		return err
	}

	// get connected directories
	connector, err := connector.NewConnector(dir)
	if err != nil {
		return err
	}
	connections := connector.Connections()

	// initialize tracker
	tracker, err := tracker.NewTracker(dir)
	if err != nil {
		return err
	}

	// update tracker at the last even if an error occurs
	defer tracker.Save()

	// delete files in connected directories
	for file := range deleted {
		for connection := range connections {
			// do not delete files in the base directory
			if connection == dir.String() {
				continue
			}

			err = types.Path(connection).Join(file).Remove()
			if err != nil {
				return err
			}
		}
		tracker.Untrack(file)
	}

	// add files in connected directories
	for file := range added {
		for connection := range connections {
			// do not add files in the base directory
			if connection == dir.String() {
				continue
			}

			err = linker.Link(file, dir, types.Path(connection))
			if err != nil {
				return err
			}
		}
		tracker.Track(file)
	}

	return nil
}
