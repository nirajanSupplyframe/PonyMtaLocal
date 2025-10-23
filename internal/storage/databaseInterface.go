package storage

import (
	"database/sql"
	"gopro/internal/dtos"
)

// / Need to do this once for DDL and to insert default values into the table.
type ExecuteInitialSqlQueries interface {
	CreateTables(db *sql.DB) error
	InsertValuesInDomainAndMTA(db *sql.DB) error
	CreateRole() error
}

//Created for the GET api. Likely will be removed or changed in need basis.
//type CheckStatusInDB interface {
//	InsertEmail(object constants.StatusEmailObject)
//}

// Creating a table to store the request that comes from post, we cannot store the message, but we can store the email address.
// and the hash value(aes256) of the message. Other than the message all other parameters can be stored.
type InsertPostfixRequest interface {
	InsertRequestData(db *sql.DB, dto dtos.PostfixRequestDTO) error
}
