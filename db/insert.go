package db

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/mgarmuno/discogsClientGo/model"
)

const (
	insertDatabase = "INSERT INTO %s (%s) VALUES (%s)"
)

// InsertUserInfo cleans and inserts the user info into the database
func InsertUserInfo(user *model.User) bool {
	db := OpenDatabase()
	tableName, fields, values := getInsertParametersFromModel(user)
	res, err := db.Exec(insertDatabase, tableName, fields, values)
	if err != nil {
		log.Println("Error inserting rows on ", tableName, ",\nfields:", fields, "\nvalues:", values, "\nerror:", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("SQL ERROR executing insert:", err)
	}
	return rowsAffected > 0
}

func getInsertParametersFromModel(model model.Model) (string, string, string) {
	tableName := reflect.TypeOf(model).Name()
	val := reflect.ValueOf(model)
	var fields []string
	var values []string
	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Name
		value := reflect.ValueOf(val.Field(i).Interface())

		fields = append(fields, fieldName)
		values = append(values, fmt.Sprint(value))
	}
	fieldsString := "(" + strings.Join(fields, ",") + ")"
	valuesString := "('" + strings.Join(values, "','") + "')"
	return tableName, fieldsString, valuesString
}
