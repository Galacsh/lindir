package cmd

import (
	"lindir/common/types"
)

type App interface {
	Init(dir types.Path) error
	Status(dir types.Path) (types.PathSet, types.PathSet, error)
	Link(fromDir, toDir types.Path) error
	Unlink(dir types.Path) error
	Push(dir types.Path, added, deleted types.PathSet) error
	Sync(dir types.Path) error
	Retrack(dir types.Path) (int, error)
}

var lindir App
