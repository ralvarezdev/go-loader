package env

import "errors"

var (
	ErrEnvironmentVariableNotFound      = "environment variable not found: %v"
	ErrFailedToLoadEnvironmentVariables = errors.New("failed to load environment variables")
	ErrInvalidDuration                  = "invalid key '%v' duration value: %v"
	ErrInvalidInteger                   = "invalid key '%v' integer value: %v"
)
