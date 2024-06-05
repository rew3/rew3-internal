package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
	// Check if the directory exists
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// The directory does not exist, so do nothing
		fmt.Println("Directory does not exist.")
		return
	} else if err != nil {
		// An error occurred, so panic
		panic(err)
	}
	// Walk through the directory and delete each file and subdirectory
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
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
 * Copy given module files to destination files.
 * Note: destination file will be created if not exists.
 * Source shoule include version too, such as github.com/projectname@version/path, it will be downloaded and copied.
 */
func CopyModuleFiles(modulePath, dstPath string) {
	cmd := exec.Command("go", "mod", "download")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	// Get the path to the module cache
	var cacheDir string
	if runtime.GOOS == "windows" {
		cacheDir = os.Getenv("USERPROFILE")
	} else {
		cacheDir = os.Getenv("HOME")
	}
	cacheDir = filepath.Join(cacheDir, "go", "pkg", "mod")
	cacheDir = cacheDir + "/" + modulePath
	CopyDirectory(cacheDir, dstPath)
}

/**
 * Copy Directory files from source directoty to destination directory.
 */
func CopyDirectory(srcPath string, dstPath string) {
	// Create the destination folder if it doesn't exist
	err := os.MkdirAll(dstPath, 0755)
	if err != nil {
		panic(err)
	}
	// Walk through the source folder and copy each file
	err = filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip this if it's the source folder itself
		if path == srcPath {
			return nil
		}
		// The relative path of the file or directory
		relPath, _ := filepath.Rel(srcPath, path)

		// The destination file or directory path
		dstFile := filepath.Join(dstPath, relPath)

		// If it's a directory, create it in the destination folder
		if info.IsDir() {
			return os.MkdirAll(dstFile, 0755)
		}
		// If it's a file, copy it to the destination folder
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
	if err != nil {
		panic(err)
	}
}

/**
 * Delete Given path file.
 */
func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}
}
