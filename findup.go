package findup

import (
	"errors"
	"os"
	"path/filepath"
)

func recursiveFind(root, fileName string) (string, error) {
	path := filepath.Join(root, fileName)
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	if root == "/" {
		return "", errors.New("File not found")
	}

	parentDir := filepath.Join(root, "..")
	return recursiveFind(parentDir, fileName)
}

// Find finds a file or a directory by walking up parent directories
func Find(fileName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return recursiveFind(dir, fileName)
}

// FindFrom finds a file or a directory  by walking up parent directories
// starting from a specified directory
func FindFrom(fileName, dir string) (string, error) {
	return recursiveFind(dir, fileName)
}
