package gaw

import (
	"os"
	"path/filepath"
)

//FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreatePath creates all directories for a file
func CreatePath(file string, mode os.FileMode) error {
	dir, _ := filepath.Split(file)
	return os.MkdirAll(dir, mode)
}
