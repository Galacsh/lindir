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

func (l lindir) Unlink(dir types.Path) error {
	return &notImplementedError{}
}
