package databaseInit

import (
	"database/sql"
	"gopro/internal/storage"
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

func (db *DatabaseStruct) InitSchema() {

	st := storage.NewExecuteSql()
	err := st.CreateTables()
	if err != nil {
		return
	}
	err3 := st.InsertValuesInDomainAndMTA(db.conn)
	if err3 != nil {
		return
	}
	err2 := st.CreateRole()
	if err2 != nil {
		return
	}

}
