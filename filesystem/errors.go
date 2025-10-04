package filesystem

import (
	"errors"
)

const (
	ErrUnableToReadFile = "unable to read file: %v"
)

var (
	ErrNilFile = errors.New("file cannot be nil")
)
