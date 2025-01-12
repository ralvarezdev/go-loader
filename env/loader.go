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
		LoadVariable(key string, dest *string) error
		LoadDurationVariable(key string, dest *time.Duration) error
		LoadSecondsVariable(key string, dest *float64) error
		LoadIntVariable(key string, dest *int) error
	}

	// DefaultLoader struct
	DefaultLoader struct {
		logger *Logger
	}
)

// NewDefaultLoader creates a new default environment variable loader
func NewDefaultLoader(loadFn func() error, logger *Logger) (
	*DefaultLoader,
	error,
) {
	// Execute the load function
	if loadFn != nil {
		if err := loadFn(); err != nil {
			return nil, ErrFailedToLoadEnvironmentVariables
		}
	}
	return &DefaultLoader{logger: logger}, nil
}

// LoadVariable load variable from environment variables
func (d *DefaultLoader) LoadVariable(key string, dest *string) error {
	// Check if the destination is nil
	if dest == nil {
		return fmt.Errorf(ErrNilDestination, key)
	}

	// Get environment variable
	variable, exists := os.LookupEnv(key)
	if !exists {
		return fmt.Errorf(ErrEnvironmentVariableNotFound, key)
	}
	*dest = variable

	// Log the environment variable
	if d.logger != nil {
		d.logger.EnvironmentVariableLoaded(key)
	}

	return nil
}

// LoadDurationVariable load duration variable from environment variables
func (d *DefaultLoader) LoadDurationVariable(
	key string,
	dest *time.Duration,
) error {
	// Check if the destination is nil
	if dest == nil {
		return fmt.Errorf(ErrNilDestination, key)
	}

	// Get environment variable
	var durationStr string
	if err := d.LoadVariable(key, &durationStr); err != nil {
		return err
	}

	// Parse the duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf(ErrInvalidDuration, key, durationStr)
	}
	*dest = duration
	return nil
}

// LoadSecondsVariable load duration variable in seconds from environment variables
func (d *DefaultLoader) LoadSecondsVariable(key string, dest *float64) error {
	// Check if the destination is nil
	if dest == nil {
		return fmt.Errorf(ErrNilDestination, key)
	}

	// Get the duration
	var duration time.Duration
	if err := d.LoadDurationVariable(key, &duration); err != nil {
		return err
	}
	*dest = duration.Seconds()
	return nil
}

// LoadIntVariable load integer variable from environment variables
func (d *DefaultLoader) LoadIntVariable(key string, dest *int) error {
	// Check if the destination is nil
	if dest == nil {
		return fmt.Errorf(ErrNilDestination, key)
	}

	// Get environment variable
	var valueStr string
	if err := d.LoadVariable(key, &valueStr); err != nil {
		return err
	}

	// Parse the integer
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fmt.Errorf(ErrInvalidInteger, key, valueStr)
	}
	*dest = value
	return nil
}
