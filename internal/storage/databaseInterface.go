package storage

import (
	constants "gopro/internal/email"
	"gopro/internal/infra/mail"
	mail2 "gopro/internal/service/mail"
)

// / Need to do this once for DDL and to insert default values into the table.
type ExecuteInitialSqlQueries interface {
	createTables() error
	insertValuesInDomainAndMTA() error
	createRole() error
}

// Creted for the GET api. Likely will be removed or changes in need basis.
type CheckStatusInDB interface {
	InsertEmail(object constants.StatusEmailObject)
}

// Using this to store the response from postfix after the post request hits the
type InsertPostfixResponse interface {
	InsertResponseData(log *mail.PostfixLogDTO) error
}

// Creating a table to store the request that comes from post, we cannot store the message but we can store the email address.
// and the hash value(aes256) of the message. Other than the message all other parameters can be stored.
type InsertPostfixRequest interface {
	InsertRequestData(log mail2.RequestDTO) error
}
