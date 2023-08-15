package util

import (
	"os"
	"path/filepath"
)

const path = ".warestack"

// ConfigFilePath returns the path for the configuration file.
func ConfigFilePath(filename string) (string, error) {
	baseDir, err := ensureBaseDirExists()
	if err != nil {
		return "", err
	}
	return filepath.Join(baseDir, filename), nil
}

// ensureBaseDirExists creates the base directory if it does not exist.
func ensureBaseDirExists() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not determine user's home directory.")
	}
	baseDir := filepath.Join(home, path)
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return baseDir, os.Mkdir(baseDir, 0755)
	}
	return baseDir, nil
}

func init() {
	_, err := ensureBaseDirExists()
	if err != nil {
		panic("Failed to ensure base directory exists: " + err.Error())
	}
}
