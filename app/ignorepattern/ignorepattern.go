package ignorepattern

import (
	"path/filepath"
	"strings"
)

type ignorePattern struct {
	negate         bool
	matchBeginning bool
	parts          []string
}

func (p ignorePattern) String() string {
	if p.negate {
		return "[!]" + strings.Join(p.parts, "/")
	} else {
		return strings.Join(p.parts, "/")
	}
}

func newIgnorePattern(pattern string) *ignorePattern {
	pattern = sanitizePattern(pattern)

	negate, pattern := handleNegate(pattern)
	matchBeginning, pattern := handleLeadingSlash(pattern)
	parts := strings.Split(pattern, "/")

	return &ignorePattern{negate, matchBeginning, parts}
}

func sanitizePattern(pattern string) string {
	pattern = strings.TrimSpace(pattern)
	pattern = filepath.ToSlash(pattern)
	return strings.TrimRight(pattern, "/")
}

func handleNegate(pattern string) (bool, string) {
	if strings.HasPrefix(pattern, "!") {
		return true, pattern[1:]
	}

	return false, pattern
}

func handleLeadingSlash(pattern string) (bool, string) {
	if strings.HasPrefix(pattern, "/") {
		return true, pattern[1:]
	}

	return false, "**/" + pattern
}

func (p ignorePattern) match(fullPath string) (bool, error) {
	if p.allAreDoubleAsterisk() {
		return true, nil
	}

	var err error
	paths := strings.Split(fullPath, "/")

	pathsIdx := 0
	partsIdx := 0

	for {
		// break if we reached the end of either the pattern or the path
		if partsIdx >= len(p.parts) || pathsIdx >= len(paths) {
			break
		}

		// current pattern and path
		pattern := p.parts[partsIdx]
		path := paths[pathsIdx]

		if pattern != "**" {
			// if pattern is not "**", use filepath.Match to match the pattern
			matched, err := filepath.Match(pattern, path)
			if err != nil || !matched {
				return false, err
			}
		} else {
			// else, use indexOf to find the next path that matches the next pattern
			partsIdx++
			if partsIdx >= len(p.parts) {
				return true, nil
			}

			pathsIdx, err = indexOf(paths, p.parts[partsIdx], pathsIdx)
			if err != nil {
				return false, err
			}

			if pathsIdx == -1 {
				return false, nil
			}
		}

		partsIdx++
		pathsIdx++
	}

	return partsIdx == len(p.parts) && pathsIdx == len(paths), nil
}

func (p ignorePattern) allAreDoubleAsterisk() bool {
	for _, part := range p.parts {
		if part != "**" {
			return false
		}
	}

	return true
}

func indexOf(paths []string, pattern string, startIdx int) (int, error) {
	for i := startIdx; i < len(paths); i++ {
		matched, err := filepath.Match(pattern, paths[i])
		if err != nil {
			return -1, err
		}

		if matched {
			return i, nil
		}
	}

	return -1, nil
}
