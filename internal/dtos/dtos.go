package dtos

import "time"

type PostfixRequestDTO struct {
	Id          string
	MessageHash string
	Timestamp   string
}

type PostfixResponseDTO struct {
	RequestId string
	Status    string
	Timestamp time.Time
}

type PostfixLogDTO struct {
	Id         string
	AccountId  string
	PostFixKey string
	CreatedOn  time.Time
	Status     string
	Reason     string
	Postfix    string
	UpdatedOn  time.Time
}
