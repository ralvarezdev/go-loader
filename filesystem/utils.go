package filesystem

import (
	"fmt"
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


// GetExecutableGoModPath get the path of the go.mod file relative to the executable (searching upwards in the directory tree)
// 
// Returns:
// 
//  - string: The path to the go.mod file
//  - error: The error if any
func GetExecutableGoModPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("could not get executable path: %w", err)
	}

   	// Start the search from the executable's directory
	dir := filepath.Dir(ex)

   	// Loop up the directory tree to find go.mod
	for {
		// Construct the potential go.mod path
		goModPath := filepath.Join(dir, "go.mod")

       // Check if go.mod exists at this location
		if _, err := os.Stat(goModPath); err == nil {
			return goModPath, nil // Found the file!
		}

       	// Move up one level
		parentDir := filepath.Dir(dir)

       	// Stop if we hit the filesystem root (no more parent directories)
		if parentDir == dir {
			return "", fmt.Errorf("go.mod not found in any parent directory relative to executable: %s", ex)
		}
		dir = parentDir
	}
}