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
 * Note: if parent directory dont exists, it will be created.
 */
func CreateFile(path string) (*os.File, error) {
	// Ensure the parent directory exists
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return nil, err
	}
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
func DeleteDirectory(dir string) error {
	// Open directory
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	// List files and directories in the directory
	files, err := d.Readdir(-1)
	if err != nil {
		return err
	}
	// Delete files and directories
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

		// If directory, recursively delete
		if file.IsDir() {
			if err := DeleteDirectory(filePath); err != nil {
				return err
			}
		} else {
			// Delete file
			if err := os.Remove(filePath); err != nil {
				return err
			}
		}
	}
	// Finally, delete the directory itself
	if err := os.Remove(dir); err != nil {
		return err
	}
	return nil
}

/**
 * Copy given module files to destination files.
 * Note: destination file will be created if not exists.
 * Source shoule include version too, such as github.com/projectname@version/path, it will be downloaded and copied.
 */
 func CopyModuleFiles(repoPath, tagVersion, schemaDir, dstPath string) {
	modulePath := repoPath + "@" + tagVersion
	println("MODULE PATH: ", modulePath)
	cmd := exec.Command("go", "mod", "download", modulePath)
	envs := append(os.Environ(), "GOPRIVATE="+repoPath)
	envs = append(envs, "GONOSUMDB="+repoPath)
	envs = append(envs, "GIT_TERMINAL_PROMPT=1")
	cmd.Env = envs
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
