package ignorepattern

import (
	"lindir/common/constants"
	"lindir/common/types"
	"path/filepath"
)

type ignorePatterns struct {
	patterns types.PathSet
}

// Return the path of the ignore patterns file
func ignorePatternsFileOf(dir types.Path) types.Path {
	return dir.Join(constants.IGNORE_PATTERNS)
}

// Return a new ignore patterns object.
// The ignore patterns object contains the patterns of the ignore patterns file.
func NewIgnorePatterns(dir types.Path) (*ignorePatterns, error) {
	file := ignorePatternsFileOf(dir)
	patterns, err := file.Read()
	if err != nil {
		return nil, err
	}

	return &ignorePatterns{patterns}, nil
}

// Check if the path matches any of the patterns
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
