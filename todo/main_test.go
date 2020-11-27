package todo

import (
	"log"
	"os"
	"testing"
)

var database *Database

const dbName = "TodoTestDB"

func TestMain(m *testing.M) {
	db, err := GetDB(dbName)
	if err != nil {
		log.Fatalf("cannot connect to database")
	}
	database = db
	os.Exit(m.Run())
}