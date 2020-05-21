package utils

import (
	"log"
	"os"
	"path/filepath"
)

// GetExecutablePath Get current executable's real path
func GetExecutablePath() (path string, dir string) {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Dir(path)

	return
}

// GetCurrentPath Get current program running directory
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
