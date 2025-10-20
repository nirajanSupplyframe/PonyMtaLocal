package databaseInit

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseStruct struct {
	conn *sql.DB
}

func NewDatabaseInit(path string) *DatabaseStruct {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	db := &DatabaseStruct{conn: conn}

	return db
}

func (db *DatabaseStruct) initSchema() {

}
