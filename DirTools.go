package gaw

import (
	"log"
	"os"
	"path"
	"strings"
)

//GetCurrentDir gets the current directory the user is is
func GetCurrentDir() string {
	exec, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return exec
}

//GetHome returns the home directory of the current user
func GetHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return home
}

//DirAbs returns the absolute path from human input (./ or ~/) and if it exists
func DirAbs(scriptPath string) (string, bool) {
	s, err := os.Stat(scriptPath)
	if err != nil || s == nil || !s.IsDir() {
		return scriptPath, false
	}
	if strings.HasPrefix(scriptPath, "/") {
		return scriptPath, true
	}

	if strings.HasPrefix(scriptPath, "./") {
		return path.Join(GetCurrentDir(), scriptPath[2:]), true
	}

	if strings.HasPrefix(scriptPath, "~/") {
		return path.Join(GetHome(), scriptPath[2:]), true
	}

	return path.Join(GetCurrentDir(), scriptPath), true
}
