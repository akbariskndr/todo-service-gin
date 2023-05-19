package mailer

import "strings"

type MailPayload struct {
	sender  *string
	to      *[]string
	cc      *[]string
	subject *string
	message *string
	html    *string
}

func (payload *MailPayload) buildFromString() string {
	return "From: " + *payload.sender + "\n"
}

func (payload *MailPayload) buildToString() string {
	return "To: " + strings.Join(*payload.to, ",") + "\n"
}

func (payload *MailPayload) buildCcString() string {
	return "Cc: " + strings.Join(*payload.cc, ",") + "\n"
}

func (payload *MailPayload) buildSubjectString() string {
	return "Subject: " + *payload.subject + "\n"
}

func (payload *MailPayload) buildMessage() string {
	var message *string
	var isHtml bool
	if payload.message == nil {
		message = payload.html
		isHtml = true
	}
	if payload.html == nil {
		message = payload.message
	}

	*message = "\n" + *message

	if isHtml {
		*message = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n" + *message
	}

	return *message
}

func (payload *MailPayload) BuildBodyString() string {
	body := payload.buildFromString() +
		payload.buildToString()

	if payload.cc != nil {
		body += payload.buildCcString()
	}

	return body + payload.buildSubjectString() + payload.buildMessage()
}
