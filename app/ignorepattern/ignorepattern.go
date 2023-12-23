package ignorepattern

import (
	"lindir/common/constants"
	"lindir/common/types"
	"path/filepath"
)

type ignorePatterns struct {
	patterns types.PathSet
}

func ignorePatternsFileOf(dir types.Path) types.Path {
	return dir.Join(constants.IGNORE_PATTERNS)
}

func NewIgnorePatterns(dir types.Path) (*ignorePatterns, error) {
	file := ignorePatternsFileOf(dir)
	patterns, err := file.Read()
	if err != nil {
		return nil, err
	}

	return &ignorePatterns{patterns}, nil
}

func (i ignorePatterns) Match(path string) (bool, error) {
	for pattern := range i.patterns {
		matched, err := filepath.Match(pattern, path)
		if err != nil {
			return false, err
		}

		if matched {
			return true, nil
		}
	}

	return false, nil
}
