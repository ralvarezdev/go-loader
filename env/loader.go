package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type (
	// Loader interface for loading environment variables
	Loader interface {
		LoadVariable(key string) (uri string, err error)
		LoadDurationVariable(key string) (duration time.Duration, err error)
		LoadSecondsVariable(key string) (seconds float64, err error)
		LoadIntVariable(key string) (value int, err error)
	}

	// DefaultLoader struct
	DefaultLoader struct{}
)

// NewDefaultLoader creates a new default environment variable loader
func NewDefaultLoader(loadFn func() error) (*DefaultLoader, error) {
	// Execute the load function
	if loadFn != nil {
		if err := loadFn(); err != nil {
			return nil, ErrFailedToLoadEnvironmentVariables
		}
	}
	return &DefaultLoader{}, nil
}

// LoadVariable load variable from environment variables
func (d *DefaultLoader) LoadVariable(key string) (uri string, err error) {
	// Get environment variable
	variable, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf(ErrEnvironmentVariableNotFound, key)
	}
	return variable, nil
}

// LoadDurationVariable load duration variable from environment variables
func (d *DefaultLoader) LoadDurationVariable(key string) (
	duration time.Duration,
	err error,
) {
	// Get environment variable
	variable, err := d.LoadVariable(key)
	if err != nil {
		return 0, err
	}

	// Parse the duration
	duration, err = time.ParseDuration(variable)
	if err != nil {
		return 0, fmt.Errorf(ErrInvalidDuration, key, variable)
	}
	return duration, nil
}

// LoadSecondsVariable load duration variable in seconds from environment variables
func (d *DefaultLoader) LoadSecondsVariable(key string) (
	seconds float64,
	err error,
) {
	// Get the duration
	duration, err := d.LoadDurationVariable(key)
	if err != nil {
		return 0, err
	}
	return duration.Seconds(), nil
}

// LoadIntVariable load integer variable from environment variables
func (d *DefaultLoader) LoadIntVariable(key string) (value int, err error) {
	// Get environment variable
	variable, err := d.LoadVariable(key)
	if err != nil {
		return 0, err
	}

	// Parse the integer
	value, err = strconv.Atoi(variable)
	if err != nil {
		return 0, fmt.Errorf(ErrInvalidInteger, key, variable)
	}
	return value, nil
}
