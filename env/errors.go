package env

import "errors"

var (
	ErrNilLoader                        = errors.New("loader cannot be nil")
	ErrNilDestination                   = "destination cannot be nil for the given key: %v"
	ErrEnvironmentVariableNotFound      = "environment variable not found: %v"
	ErrFailedToLoadEnvironmentVariables = errors.New("failed to load environment variables")
	ErrInvalidDuration                  = "invalid key '%v' duration value: %v"
	ErrInvalidInteger                   = "invalid key '%v' integer value: %v"
)
