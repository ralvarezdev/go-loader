package env

import (
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
)
