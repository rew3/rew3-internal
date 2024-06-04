package utils

import (
	"fmt"
	"os"
)

/**
 * Create a new file for given file path.
 */
func CreateFile(path string) (*os.File, error) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil, err
	}
	return outputFile, nil
}

/**
 * Write a data into given filepath file (append)
 */
func WriteToFile(filePath, data string) error {
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

/**
 * Make a new directory for given path.
 */
func MakeDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, os.ModePerm)
	if err != nil {
		fmt.Println("Unable to create directory: ", err)
	}
	fmt.Println("Directory created:", directoryPath)
	return nil
}

/**
 * Check if given file with name/path already exists or not.
 */
func IsFileAlreadyExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File '%s' does not exist.\n", filename)
			return false
		} else {
			fmt.Printf("Error checking file: %v\n", err)
			return false
		}
	} else {
		fmt.Printf("File '%s' already exists.\n", filename)
		return true
	}
}
