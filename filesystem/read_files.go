package filesystem

import (
	"os"
)

// ReadFile reads the file from the given path
func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, UnableToReadFileError
	}
	return file, nil
}
