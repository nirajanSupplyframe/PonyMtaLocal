package mail

type Sender interface {
	SendMail(id, to, subject, body string) (string, error)
}
