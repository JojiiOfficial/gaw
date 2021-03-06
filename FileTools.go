package gaw

import (
	"crypto/md5"
	"encoding/hex"
	"io"
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

// FileFromPath returns file from given path
func FileFromPath(path string) string {
	_, file := filepath.Split(path)
	return file
}

// PathFromFilepath returns a path from a filepath
func PathFromFilepath(fp string) string {
	path, _ := filepath.Split(fp)
	return path
}

// GetFileMD5 get md5 hash of file
func GetFileMD5(filePath string) (string, error) {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buff := make([]byte, 1024*10)
	hash := md5.New()

	// Copy file into hash
	if _, err := io.CopyBuffer(hash, file, buff); err != nil {
		return "", err
	}

	// Get Sum and encode it to a readable hex string
	return hex.EncodeToString(hash.Sum(nil)[:16]), nil
}
