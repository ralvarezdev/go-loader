package env

import (
	"fmt"
	"os"
	"time"
)

// LoadVariable load variable from environment variables
func LoadVariable(key string) (uri string, err error) {
	// Get environment variable
	variable, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf(ErrEnvironmentVariableNotFound, key)
	}
	return variable, nil
}

// LoadDurationVariable load duration variable from environment variables
func LoadDurationVariable(key string) (duration time.Duration, err error) {
	// Get environment variable
	variable, err := LoadVariable(key)
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
func LoadSecondsVariable(key string) (seconds float64, err error) {
	// Get the duration
	duration, err := LoadDurationVariable(key)
	if err != nil {
		return 0, err
	}
	return duration.Seconds(), nil
}
