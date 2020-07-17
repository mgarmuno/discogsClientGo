package file

import (
	"io/ioutil"
	"log"
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

// GetToken returns the token saved in the file
func GetToken() string {
	token, err := ioutil.ReadFile(tokenFileName)
	if err != nil {
		log.Println("Error reading token:", err)
	}

	return string(token)
}

// SaveToken saves the token into a txt file
func SaveToken(token string) bool {
	err := os.Remove(tokenFileName)
	if err != nil {
		log.Println("Error removing token file:", err)
	}

	file, err := os.Create(tokenFileName)
	if err != nil {
		log.Println("Error creating file to save the token:", err)
		return false
	}

	_, err = file.WriteString(token)
	if err != nil {
		file.Close()
		log.Println("Error writing string into token file:", err)
		return false
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing token file:", err)
		return false
	}
	return true
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
