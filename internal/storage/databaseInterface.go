package storage

import (
	constants "gopro/internal/email"
)

type executeSqlQueries interface {
	createTables() error
	insertValuesInDomainAndMTA() error
	createRole() error
}
type CheckStatusInDB interface {
	InsertEmail(object constants.EmailObject)
}
