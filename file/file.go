package file

import (
	"os"
)

const (
	tokenFileName = "token.txt"
)

// CheckToken checks if the file with the token exists.
func CheckToken() bool {
	if !checkFileExists() {
		return false
	}
	return checkToken()
}

func checkFileExists() bool {
	if _, err := os.Stat(tokenFileName); err == nil {
		return true
	}
	return false
}

func checkToken() bool {
	fi, err := os.Stat(tokenFileName)
	if err != nil {
		return false
	}
	return fi.Size() > 0
}
