package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/google/uuid"
)

type PostfixSender struct {
	addr string
	from string
}

// NewPostfixSender This is an implementation of the interface created in sendMailInterface.go
func NewPostfixSender(addr string) Sender {
	return &PostfixSender{
		addr: addr,
		from: "noreply@example.local",
	}
}

// SendMail : function is used to send mail using postfix using the parameters send from the post request by the client.
func (p *PostfixSender) SendMail(to, subject, body string) (string, error) {
	id := uuid.New().String()
	messageID := fmt.Sprintf("<%s@example.com>", id)
	xInternal := id

	var b bytes.Buffer
	b.WriteString("From: " + p.from)
	b.WriteString("To: " + to)
	b.WriteString("Subject: " + subject)
	b.WriteString("MessageID: " + messageID)
	b.WriteString("X-Internal-ID: " + xInternal)
	//post fix email formating requires a line gap between headers and message body
	b.WriteString("\r\n")
	b.WriteString(body)

	//msg := []byte("To: " + to + "\r\n" +
	//	"Subject: " + subject + "" +
	//	"\r\n\r\n" + body)

	//err := smtp.SendMail("localhost:25", nil, from, []string{to}, msg)

	if err := smtp.SendMail(p.addr, nil, p.from, []string{to}, b.Bytes()); err != nil {
		return "", err
	}
	return id, nil
}

// Helper: to get id from message id . Not used but may come in handy
func requestIDFromMessageID(msgid string) string {
	msgid = strings.TrimSpace(msgid)
	msgid = strings.TrimPrefix(msgid, "<")
	msgid = strings.TrimSuffix(msgid, ">")
	parts := strings.SplitN(msgid, "@", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
