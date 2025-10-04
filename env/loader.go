package env

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type (
	// DefaultLoader struct
	DefaultLoader struct {
		logger *slog.Logger
	}
)

// NewDefaultLoader creates a new default environment variable loader
//
// Parameters:
//
//   - loadFn: function to load the environment variables (e.g., from a file)
//   - logger: logger instance for logging
//
// Returns:
//
//   - *DefaultLoader: instance of DefaultLoader
//   - error: error if any occurred during loading
func NewDefaultLoader(loadFn func() error, logger *slog.Logger) (
	*DefaultLoader,
	error,
) {
	// Execute the load function
	if loadFn != nil {
		if err := loadFn(); err != nil {
			return nil, ErrFailedToLoadEnvironmentVariables
		}
	}

	// Prepare the logger
	if logger != nil {
		logger = logger.With("component", "env_loader")
	}
	return &DefaultLoader{logger}, nil
}

// LoadVariable load variable from environment variables
//
// Parameters:
//
//   - key: environment variable key
//   - dest: pointer to the destination string where the value will be stored
//
// Returns:
//
//   - error: error if any occurred during loading or parsing
func (d DefaultLoader) LoadVariable(key string, dest *string) error {
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
	EnvironmentVariableLoaded(d.logger, key)

	return nil
}

// LoadDurationVariable load duration variable from environment variables
//
// Parameters:
//
//   - key: environment variable key
//   - dest: pointer to the destination time.Duration where the value will be stored
//
// Returns:
//
//   - error: error if any occurred during loading or parsing
func (d DefaultLoader) LoadDurationVariable(
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
//
// Parameters:
//
//   - key: environment variable key
//   - dest: pointer to the destination float64 where the value in seconds will be stored
//
// Returns:
//
//   - error: error if any occurred during loading or parsing
func (d DefaultLoader) LoadSecondsVariable(key string, dest *float64) error {
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
//
// Parameters:
//
//   - key: environment variable key
//   - dest: pointer to the destination int where the value will be stored
//
// Returns:
//
//   - error: error if any occurred during loading or parsing
func (d DefaultLoader) LoadIntVariable(key string, dest *int) error {
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
