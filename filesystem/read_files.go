package filesystem

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadFile reads the file from the given path
func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToReadFile, err)
	}
	return file, nil
}

// OpenFile opens the file from the given path
func OpenFile(path string) (*os.File, error) {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return file, nil
}

// CloseFile closes the file
func CloseFile(file *os.File) error {
	if file != nil {
		err := file.Close()
		if err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}
	}
	return nil
}

// ReadCSVFile reads a CSV file and returns a slice of string slices
func ReadCSVFile(file *os.File, readHeaders bool) (
	*[][]string,
	*[]string,
	error,
) {
	// Check if the file is nil
	if file == nil {
		return nil, nil, ErrNilFile
	}

	// Defer closing the file
	defer func(file *os.File) {
		if err := CloseFile(file); err != nil {
			fmt.Println(err)
		}
	}(file)

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header line
	var headers []string
	var err error
	if readHeaders {
		headers, err = reader.Read()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read header line: %w", err)
		}
	}

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	if readHeaders {
		return &records, &headers, nil
	}
	return &records, nil, nil
}
