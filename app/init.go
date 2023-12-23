package app

import "lindir/common/types"

func (l lindir) Init(dir types.Path) error {
	notInitialized, err := isNotInitialized(dir)
	if err != nil {
		return err
	}

	if notInitialized {
		if err = createAppDir(dir); err != nil {
			return err
		}
		if err = createTrackerFile(dir); err != nil {
			return err
		}
		if err = createConnectorFile(dir); err != nil {
			return err
		}
	} else {
		return alreadyInitializedError{dir}
	}

	return nil
}

func isNotInitialized(dir types.Path) (bool, error) {
	return false, &notImplementedError{}
}
