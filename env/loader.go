package env

import (
	"fmt"
	"os"
)

// LoadVariable load variable from environment variables
func LoadVariable(key string) (uri string, err error) {
	// Get environment variable
	variable, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf(EnvironmentVariableNotFoundError, key)
	}
	return variable, nil
}
