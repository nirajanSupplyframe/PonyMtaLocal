package storage

//
//import (
//	"database/sql"
//	"gopro/internal/infra/mail"
//)
//
//type ExecuteSql struct{}
//
//func NewExecuteInsertPostfix() ExecuteInitialSqlQueries {
//	return &ExecuteSql{}
//}
//
//func NewExecuteSql() InsertPostfixResponse {
//	return &ExecuteSql{}
//}
//
//func (e *ExecuteSql) createTables() error {
//	return nil
//}
//
//func (e *ExecuteSql) insertValuesInDomainAndMTA() error {
//	return nil
//}
//
//func (e *ExecuteSql) createRole() error {
//	return nil
//}
//
//func (e *ExecuteSql) InsertResponseData(log *mail.PostfixLogDTO) error {
//	query := "\nINSERT INTO postfix_logs \n    (id, account_id,postfix_key,created_on,status,reason,postfix_log,updated_on)\n    VALUES (?,?,?,?,?,?,?,?)\n    \t\t\n\t"
//	db := sql.DB{}
//	_, err := db.Exec(query,
//
//		log.ProcessID,
//		log.ProcessID, //must be application id but putting process id which is postfix id for the meanwhile
//		log.ProcessID,
//		log.Timestamp,
//		log.Status,
//		log.Message,
//		log.Raw,
//		log.Timestamp)
//
//	return err
//}
//
//func (e *ExecuteSql) InsertPostfixLog(log *mail.PostfixLogDTO) error {
// return nil
//}
