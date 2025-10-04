package env

import (
	"errors"
)

const (
	ErrNilDestination              = "destination cannot be nil for the given key: %v"
	ErrEnvironmentVariableNotFound = "environment variable not found: %v"
	ErrInvalidDuration             = "invalid key '%v' duration value: %v"
	ErrInvalidInteger              = "invalid key '%v' integer value: %v"
)

var (
	ErrNilLoader                        = errors.New("loader cannot be nil")
	ErrFailedToLoadEnvironmentVariables = errors.New("failed to load environment variables")
)
