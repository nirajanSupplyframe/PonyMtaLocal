package databaseInit

import (
	"database/sql"
	"gopro/internal/dtos"
)

type DBHelper struct {
	db *sql.DB
}

func (h *DBHelper) InsertIntoEmailDatabase(dtos.PostfixLogDTO) bool {
	query := "Insert INTO main.email_status ()"
	return true
}
