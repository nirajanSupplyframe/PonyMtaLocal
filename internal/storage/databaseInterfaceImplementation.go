package storage

import (
	"database/sql"
	"errors"
	"gopro/internal/dtos"
	"log"
	"os"
	"time"
)

type ExecuteSql struct {
	db *sql.DB
}

func NewExecuteSql() *ExecuteSql {
	return &ExecuteSql{}
}

func (e *ExecuteSql) CreateTables() error {

	schema, err := os.ReadFile("./internal/scripts/schemas.sql")
	if err != nil {
		log.Fatalf("Unable to read schema file %v", err)
	}

	_, err = e.db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Unable to execute schema %v", err)
	}

	log.Println("Schema executed successfully!")
	return nil
}

func (e *ExecuteSql) InsertValuesInDomainAndMTA() error {
	schema, err := os.ReadFile("./internal/scripts/defaults.sql")
	if err != nil {
		log.Fatalf("Unable to read schema file %v", err)
	}
	_, err = e.db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Unable to execute schema %v", err)
	}
	return nil
}

func (e *ExecuteSql) CreateRole() error {
	return nil
}

func (e *ExecuteSql) InsertRequestData(log dtos.PostfixRequestDTO) error {
	//query := "\nINSERT INTO postfix_logs \n    (id, account_id,postfix_key,created_on,status,reason,postfix_log,updated_on)\n    VALUES (?,?,?,?,?,?,?,?)\n    \t\t\n\t"
	//db := sql.DB{}
	//_, err := db.Exec(query,

	query := "\nINSERT INTO post_request \n (id,message,created_on)\n VALUES (?,?,?)\n \t\t\n\t"
	_, err := e.db.Exec(query,
		log.Id,
		log.MessageHash,
		time.Now())

	return err
}

func (e *ExecuteSql) CheckReqIdInDatabase(reqId string) error {
	query := "\nSELECT id FROM email_status WHERE id = ?"

	var foundId string
	err := e.db.QueryRow(query, reqId).Scan(&foundId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	return nil
}

func (e *ExecuteSql) InsertPostfixLog(log dtos.PostfixLogDTO) error {
	return nil
}
