package mail

import (
	"fmt"
	"net/smtp"
)

type PostfixSender struct {
}

// NewPostfixSender This is an implementation of the interface created in sendMailInterface.go
func NewPostfixSender() Sender {
	return &PostfixSender{}
}

// SendMail : function is used to send mail using postfix using the parameters send from the post request by the client.
func (p *PostfixSender) SendMail(to, subject, body string) error {
	from := "abc@def.com"
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "" +
		"\r\n\r\n" + body)

	err := smtp.SendMail("localhost:25", nil, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}
	return nil
}
