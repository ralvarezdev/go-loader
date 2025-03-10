package filesystem

import (
	"errors"
)

var (
	ErrUnableToReadFile = "unable to read file: %v"
	ErrNilFile          = errors.New("file cannot be nil")
)
