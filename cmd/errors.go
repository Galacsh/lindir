package cmd

import (
	"fmt"
)

type cannotGetDirectory struct {
	cmd string
	err error
}

func (e cannotGetDirectory) Error() string {
	return fmt.Sprintf("failed to get directory for '%s': %s", e.cmd, e.err.Error())
}

type initError struct {
	dir string
	err error
}

func (e initError) Error() string {
	return fmt.Sprintf("failed to initialize '%s': %s", e.dir, e.err.Error())
}

type linkError struct {
	fromDir string
	toDir   string
	err     error
}

func (e linkError) Error() string {
	return fmt.Sprintf("failed to link '%s' to '%s': %s", e.fromDir, e.toDir, e.err.Error())
}

type statusError struct {
	dir string
	err error
}

func (e statusError) Error() string {
	return fmt.Sprintf("failed to get status of '%s': %s", e.dir, e.err.Error())
}

type pushError struct {
	dir string
	err error
}

func (e pushError) Error() string {
	return fmt.Sprintf("failed to push '%s': %s", e.dir, e.err.Error())
}

type syncError struct {
	dir string
	err error
}

func (e syncError) Error() string {
	return fmt.Sprintf("failed to sync '%s': %s", e.dir, e.err.Error())
}

type unlinkError struct {
	dir string
	err error
}

func (e unlinkError) Error() string {
	return fmt.Sprintf("failed to unlink '%s': %s", e.dir, e.err.Error())
}

type retrackError struct {
	dir string
	err error
}

func (e retrackError) Error() string {
	return fmt.Sprintf("failed to retrack '%s': %s", e.dir, e.err.Error())
}
