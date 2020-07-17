package db

import (
	"log"

	"github.com/mgarmuno/discogsClientGo/model"
)

const (
	simpleSelect = "SELECT * FROM %d"
	user         = "User"
)

// GetUserInfo retrieves the user information from the DB
func GetUserInfo() model.User {
	db := OpenDatabase()
	rows, err := db.Query(simpleSelect, user)
	var user model.User
	if err != nil {
		log.Println("Error executing select ", user, ":", err)
		return user
	}
	for rows.Next() {
		err = rows.Scan(user)
	}
	rows.Close()
	return user
}

// GetLanguage returns the language setted in the database.
func GetLanguage() string {
	return "en"
}
