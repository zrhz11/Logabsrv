package internal

import (
	"testing"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

func loadenv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Fprintln(os.Stderr ,"Error loading .env file")
	}
}

func resetDB() {
	DBPath := os.Getenv("DBPath")

	err := os.Remove(DBPath)
	if err != nil {
		fmt.Fprintf(os.Stderr ,"Failed to delete DB file: %v", err)
	}

	file, err := os.Create(DBPath)
	if err != nil {
		fmt.Fprintf(os.Stderr ,"Failed to create DBfile: %v", err)
	}
	file.Close()
}

func init() {
	loadenv()
	resetDB()
}


func TestSQLInit(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Fatalf("error init DB: %v" , err)
	}
}