package mail

type Sender interface {
	SendMail(from, to, subject, body string) error
}
