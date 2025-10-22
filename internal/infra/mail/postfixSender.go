package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
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
		from: "nirajanchapagain@SF-NCHAPAGA.localdomain",
	}
}

// SendMail : function is used to send mail using postfix using the parameters send from the post request by the client.
func (p *PostfixSender) SendMail(to, subject, body string) (string, error) {
	id := uuid.New().String()
	host, _ := os.Hostname()
	println(host)
	println("Real id created at sendmail function :"+id, host+".localdomain")
	messageID := fmt.Sprintf("<%s@%s>", id, host+".localdomain")
	xInternal := id
	println("Message id after change :" + messageID)

	var b bytes.Buffer
	b.WriteString("From: " + p.from + "\r\n")
	b.WriteString("To: " + to + "\r\n")
	b.WriteString("Subject: " + subject + "\r\n")
	b.WriteString("Message-ID: " + messageID + "\r\n")
	b.WriteString("X-Internal-ID: " + xInternal + "\r\n")
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
