package app

import (
	"lindir/common/constants"
	"lindir/common/types"
)

func createAppDir(dir types.Path) error {
	return appDirOf(dir).CreateDir()
}

func appDirOf(dir types.Path) types.Path {
	return dir.Join(constants.APP_DIR)
}
