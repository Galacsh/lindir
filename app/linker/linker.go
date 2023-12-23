package linker

import (
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/types"
	"os"
)

func LinkAppDir(from, to types.Path) error {
	trackerFile := tracker.TrackerFileOf(from)
	err := Link(trackerFile.String(), from, to)
	if err != nil {
		return err
	}

	connectorFile := connector.ConnectorFileOf(to)
	err = Link(connectorFile.String(), from, to)
	if err != nil {
		return err
	}

	return nil
}

func LinkTrackedFiles(from, to types.Path) error {
	tracker, err := tracker.NewTracker(from)
	if err != nil {
		return err
	}

	for file := range tracker.TrackingFiles() {
		err = Link(file, from, to)
		if err != nil {
			return err
		}
	}

	return nil
}

func Link(relPath string, from, to types.Path) error {
	fromFile := from.Join(relPath)
	fromStat, err := os.Stat(fromFile.String())
	if err != nil {
		return err
	}

	toFile := to.Join(relPath)
	toStat, err := os.Stat(toFile.String())

	// create link if file does not exist
	if os.IsNotExist(err) {
		err = toFile.Dir().CreateDir()
		if err != nil {
			return err
		}
		return os.Link(fromFile.String(), toFile.String())
	} else if err != nil {
		return err
	}

	if os.SameFile(fromStat, toStat) {
		// do nothing if file is already linked
		return nil
	} else {
		// return error if different file with same name exists
		return fileWithSameNameExistsError{toFile}
	}

}
