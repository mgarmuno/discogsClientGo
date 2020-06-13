package db

import (
	"database/sql"
	"log"
	"os"
)

const (
	driver   = "sqlite3"
	fileName = "discogsDatabase.sql"
	dbDir    = "db/"
)

// CheckDatabase checks if the database file exists
func CheckDatabase() {
	exists, err := checkDatabaseExists()
	if err != nil {
		log.Fatal("Problem checking database file:", err)
	}
	if !exists {
		createDatabase()
	}
}

func checkDatabaseExists() (bool, error) {
	_, err := os.Stat("./" + dbDir + fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDatabase() {
	db, err := sql.Open(driver, "./"+dbDir+fileName)
	if err != nil {
		log.Fatal("Error creating database file", dbDir+fileName, ":", err)
	}
}
