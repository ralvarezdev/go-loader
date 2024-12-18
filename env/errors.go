package env

import "errors"

var (
	EnvironmentVariableNotFoundError      = "environment variable not found: %v"
	FailedToLoadEnvironmentVariablesError = errors.New("failed to load environment variables")
)
