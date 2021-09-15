package email

import (
	"fmt"
	"net/smtp"
)

type FakeGmail struct {
	config *SMTPConfig
}

var receive_email []string

func InitFakeGmail(config *SMTPConfig, test_receive_email string) {
	Emailer = FakeGmail{
		config: config,
	}
	receive_email = []string{test_receive_email}
}

func (gmail FakeGmail) SendPlainEmail(to []string, subject string, body string) error {
	emailAuth := smtp.PlainAuth("me", gmail.config.From, gmail.config.Password, gmail.config.Host)
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", gmail.config.Host, gmail.config.Port)

	if err := smtp.SendMail(addr, emailAuth, gmail.config.From, receive_email, msg); err != nil {
		return err
	}
	return nil
}

func (gmail FakeGmail) SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error {
	emailAuth := smtp.PlainAuth("me", gmail.config.From, gmail.config.Password, gmail.config.Host)

	body, err := renderHTML(data, layout...)

	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", gmail.config.Host, gmail.config.Port)

	if err := smtp.SendMail(addr, emailAuth, gmail.config.From, receive_email, msg); err != nil {
		return err
	}
	return nil

}
