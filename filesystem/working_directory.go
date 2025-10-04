package filesystem

import (
	"os"
	"path/filepath"
)

// GetCurrentDirectory gets the current directory of the executable
//
// Returns:
//
//   - string: the current directory path
//   - error: error if any occurred while retrieving the directory
func GetCurrentDirectory() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(execPath), nil
}
