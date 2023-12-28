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
	defaultPatterns := types.Paths{}
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

	patterns, err := file.ReadPaths()
	if err != nil {
		return nil, err
	}

	compiled, err := convertPatterns(defaultPatterns.Concat(patterns))
	if err != nil {
		return nil, err
	}
	return &ignorePatterns{compiled}, nil
}

func convertPatterns(patterns types.Paths) ([]ignorePattern, error) {
	converted := make([]ignorePattern, 0, len(patterns))
	for _, pattern := range patterns {
		converted = append(converted, *newIgnorePattern(pattern.String()))
	}

	return converted, nil
}

// Return the path of the ignore patterns file
func ignorePatternsFileOf(dir types.Path) types.Path {
	return dir.Join(constants.IGNORE_PATTERNS)
}

// Check if the path matches any of the patterns
func (i ignorePatterns) ShouldIgnore(path string) (bool, error) {
	ignore := false

	for _, pattern := range i.patterns {
		matched, err := pattern.match(path)

		if err != nil {
			return false, err
		}

		// on negation, excluded by previous patterns should be included again
		if matched && pattern.negate {
			ignore = false
		} else {
			ignore = ignore || matched
		}
	}

	return ignore, nil
}
