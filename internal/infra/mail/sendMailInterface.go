package mail

type Sender interface {
	SendMail(to, subject, body string) (string, error)
}
