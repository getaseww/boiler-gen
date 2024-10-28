package helpers

import (
	"fmt"
	"os"
)

// Function to create a file with given content
func CreateFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// Function to check and create a directory if it doesn't exist
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("Created directory: %s\n", dir)
	}
	return nil
}
