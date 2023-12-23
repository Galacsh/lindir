package app

import "lindir/common/types"

type lindir struct{}

func New() lindir {
	return lindir{}
}

type notImplementedError struct{}

func (e notImplementedError) Error() string {
	return "not implemented"
}

func (l lindir) Status(dir types.Path) (types.PathSet, types.PathSet, error) {
	return nil, nil, &notImplementedError{}
}

func (l lindir) Link(fromDir, toDir types.Path) error {
	return &notImplementedError{}
}

func (l lindir) Unlink(dir types.Path) error {
	return &notImplementedError{}
}

func (l lindir) Push(dir types.Path) error {
	return &notImplementedError{}
}

func (l lindir) Sync(dir types.Path) error {
	return &notImplementedError{}
}
