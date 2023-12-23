package app

import "lindir/common/types"

func createConnector(dir types.Path) error {
	return &notImplementedError{}
}

func initConnector(dir types.Path) error {
	return &notImplementedError{}
}

func areConnected(from types.Path, to types.Path) (bool, error) {
	return false, &notImplementedError{}
}

func connect(from types.Path, to types.Path) error {
	return &notImplementedError{}
}
