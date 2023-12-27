package ignorepattern

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type ignorePatterns struct {
	// compiled glob patterns
	patterns []ignorePattern
}

// Return a new ignore patterns object.
// The ignore patterns object contains the patterns of the ignore patterns file.
func NewIgnorePatterns(dir types.Path) (*ignorePatterns, error) {
	defaultPatterns := make(types.PathSet)
	defaultPatterns.Add(constants.APP_DIR)

	file := ignorePatternsFileOf(dir)
	notExists, err := file.NotExists()
	if err != nil {
		return nil, err
	}

	if notExists {
		converted, err := convertPatterns(defaultPatterns)
		if err != nil {
			return nil, err
		}
		return &ignorePatterns{converted}, nil
	}

	patterns, err := file.Read()
	if err != nil {
		return nil, err
	}

	compiled, err := convertPatterns(defaultPatterns.Union(patterns))
	if err != nil {
		return nil, err
	}
	return &ignorePatterns{compiled}, nil
}

func convertPatterns(patterns types.PathSet) ([]ignorePattern, error) {
	converted := make([]ignorePattern, 0, len(patterns))
	for pattern := range patterns {
		converted = append(converted, *newIgnorePattern(pattern))
	}

	return converted, nil
}

// Return the path of the ignore patterns file
func ignorePatternsFileOf(dir types.Path) types.Path {
	return dir.Join(constants.IGNORE_PATTERNS)
}

// Check if the path matches any of the patterns
func (i ignorePatterns) Match(path string) (bool, error) {
	for _, pattern := range i.patterns {
		matched, err := pattern.match(path)
		if err != nil || matched {
			return matched, err
		}
	}

	return false, nil
}
