package utils

import (
	"fmt"
	"io"
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

// IndexOfStringList Find the position of a string in a string list
// return -1 if not found in string list
func IndexOfStringList(value string, list []string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == value {
			return i
		}
	}

	return -1
}

// Copy file from src path to dist path
func Copy(src, dest string) (written int64, err error) {
	srcFile, err := os.Open(src)

	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()

	dstFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
