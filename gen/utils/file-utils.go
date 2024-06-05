package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
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

/**
 * Delete given path directory.
 */
func DeleteDirectory(dirPath string) {
	// Walk through the directory and delete each file and subdirectory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip this if it's the directory itself
		if path == dirPath {
			return nil
		}
		// If it's a directory, delete it and its contents
		if info.IsDir() {
			return os.RemoveAll(path)
		}
		// If it's a file, delete it
		return os.Remove(path)
	})
	if err != nil {
		panic(err)
	}
	// Delete the directory itself
	err = os.Remove(dirPath)
	if err != nil {
		panic(err)
	}
}

/**
 * Copy all content of source path directory to given destination path directory.
 */
func CopyModuleFiles(sourcePath, dstPath string) error {
	// Download and cache the module files
	cmd := exec.Command("go", "mod", "download", sourcePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	// Get the path to the module cache
	cacheDir, _ := os.UserCacheDir()
	cacheDir = filepath.Join(cacheDir, "go-build")
	// Find the module directory in the cache
	moduleDir := ""
	err := filepath.Walk(cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == sourcePath {
			moduleDir = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return err
	}
	if moduleDir == "" {
		return os.ErrNotExist
	}
	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(dstPath, 0755)
	if err != nil {
		return err
	}
	// Walk through the module directory and copy each file
	return filepath.Walk(moduleDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip this if it's the module directory itself
		if path == moduleDir {
			return nil
		}
		// The relative path of the file or directory
		relPath, _ := filepath.Rel(moduleDir, path)
		// The destination file or directory path
		dstFile := filepath.Join(dstPath, relPath)
		// If it's a directory, create it in the destination directory
		if info.IsDir() {
			return os.MkdirAll(dstFile, 0755)
		}
		// If it's a file, copy it to the destination directory
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		file, err := os.Create(dstFile)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, srcFile)
		return err
	})
}
