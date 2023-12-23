package linker

import (
	"io"
	"lindir/app/connector"
	"lindir/app/tracker"
	"lindir/common/types"
	"os"
)

// Hard link files inside 'from/.lindir' directory to 'to/.lindir' directory
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

// Hard link tracked files inside 'from' directory to 'to' directory
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

// Hard link 'from' file to 'to' file
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

// Turn hard link into copy of the file
func HardLinkToCopy(file types.Path) error {
	temp, err := copyTemp(file)
	if err != nil {
		return err
	}

	// Remove the original file
	err = os.Remove(file.String())
	if err != nil {
		return err
	}

	// Rename the copy to the original file name
	return os.Rename(temp, file.String())
}

// Copy the file to a temporary file
func copyTemp(file types.Path) (string, error) {
	from, err := os.Open(file.String())
	if err != nil {
		return "", err
	}
	defer from.Close()

	temp, err := os.CreateTemp("", "*")
	if err != nil {
		return "", err
	}
	defer temp.Close()

	_, err = io.Copy(temp, from)
	if err != nil {
		return "", err
	}

	return temp.Name(), nil
}
