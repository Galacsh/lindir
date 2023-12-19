package cmd

import "lindir/common/types"

type App interface {
	Init(dir types.Path) error
	Status(dir types.Path) (types.PathSet, types.PathSet, error)
	Link(fromDir, toDir types.Path) error
	Unlink(dir types.Path) error
	Push(dir types.Path) error
	Sync(dir types.Path) error
}

var lindir App
