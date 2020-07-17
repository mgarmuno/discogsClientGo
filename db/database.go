package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mgarmuno/discogsClientGo/model"
)

const (
	driver      = "sqlite3"
	fileName    = "discogsDatabase.db"
	create      = "CREATE TABLE IF NOT EXISTS "
	vinyl       = "vinyl"
	vinylFields = "id integer primary key, title text, artist integer, released integer, genre text, style text"
	artist      = "artist"
	label       = "label"
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

// OpenDatabase creates a DB connection
func OpenDatabase() *sql.DB {
	db, err := sql.Open(driver, "./"+fileName)
	if err != nil {
		log.Fatal("Error opening database file", fileName, ":", err)
	}
	return db
}

func checkDatabaseExists() (bool, error) {
	_, err := os.Stat("./" + fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDatabase() {
	fmt.Println("Creating database...")
	db := OpenDatabase()
	fmt.Println("Database created")
	createDBTables(db)
}

func createDBTables(db *sql.DB) {
	createTable(db, model.Vinyl{})
	createTable(db, model.Artist{})
	createTable(db, model.Uris{})
	createTable(db, model.User{})
}

func createTable(db *sql.DB, model model.Model) {
	tableName := reflect.TypeOf(model).Name()
	val := reflect.ValueOf(model)
	var fields []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag
		fieldName := field.Name
		fieldTag := tag.Get("sql")
		fields = append(fields, fmt.Sprintln(fieldName, fieldTag))
	}
	fieldsString := "(" + strings.Join(fields, ",") + ")"
	_, err := db.Exec(fmt.Sprint(create, tableName, fieldsString))
	if err != nil {
		log.Fatal("Error creating table", tableName, ":", err)
	}
	fmt.Println("Table", tableName, "created")
}
