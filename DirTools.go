package goawesomehelper

import (
	"log"
	"os"
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
