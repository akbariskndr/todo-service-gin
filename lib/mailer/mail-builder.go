package mailer

import (
	"fmt"
	"net/smtp"
)

type MailBuilder struct {
	config *MailConfig
	mail   *MailPayload
}

func (mailer *MailBuilder) To(val *[]string) *MailBuilder {
	mailer.mail.to = val

	return mailer
}

func (mailer *MailBuilder) Cc(val *[]string) *MailBuilder {
	mailer.mail.cc = val

	return mailer
}

func (mailer *MailBuilder) Subject(val *string) *MailBuilder {
	mailer.mail.subject = val

	return mailer
}

func (mailer *MailBuilder) Message(val *string) *MailBuilder {
	mailer.mail.message = val

	return mailer
}

func (mailer *MailBuilder) Html(html string) *MailBuilder {
	mailer.mail.html = &html

	return mailer
}

func (mailer *MailBuilder) Send() {
	go func() {
		body := mailer.mail.BuildBodyString()

		auth := smtp.PlainAuth("", mailer.config.username, mailer.config.password, mailer.config.host)
		smtpAddr := fmt.Sprintf("%s:%d", mailer.config.host, mailer.config.port)

		cc := []string{}
		if mailer.mail.cc != nil {
			cc = *mailer.mail.cc
		}
		err := smtp.SendMail(smtpAddr, auth, mailer.config.sender, append(*mailer.mail.to, cc...), []byte(body))
		if err != nil {
			panic(err.Error())
		}
	}()
}
