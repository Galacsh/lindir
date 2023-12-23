package app

import (
	"lindir/app/check"
	"lindir/app/connector"
	"lindir/app/linker"
	"lindir/app/tracker"
	"lindir/common/constants"
	"lindir/common/types"
)

// Unlink 'dir' directory
func (l lindir) Unlink(dir types.Path) error {
	err := check.ErrIfNotInitialized(dir)
	if err != nil {
		return err
	}

	// disconnect directory
	connector, err := connector.NewConnector(dir)
	if err != nil {
		return err
	}

	connector.Disconnect(dir)
	connector.Save()

	// change hard links to whole new copies
	tracker, err := tracker.NewTracker(dir)
	if err != nil {
		return err
	}

	files := tracker.TrackingFiles()
	for file := range files {
		err = linker.HardLinkToCopy(dir.Join(file))
		if err != nil {
			return err
		}
	}

	// remove app directory
	err = dir.Join(constants.APP_DIR).Remove()
	if err != nil {
		return err
	}

	return nil
}
