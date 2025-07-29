package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() (error) {

	DBPath := os.Getenv("DBPath")

	var err error
	db, err = sql.Open("sqlite3", DBPath)
	if err != nil {
		fmt.Fprintf(os.Stderr , "error connecting to db: %v" , err)
	}

	DBSchemaPath := os.Getenv("DBSchemaPath")

	if DBSchemaPath == "" {
		return nil
	}

	sqlBytes, err := os.ReadFile(DBSchemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading DB schema: %v" , err)
	}
	sqlString := string(sqlBytes)

	_, err = db.Exec(sqlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error execing DB schema: %v" , err)
	}

	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}