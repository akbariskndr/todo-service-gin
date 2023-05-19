package mailer

import (
	"fmt"
	"os"
	"strconv"
)

type MissingConfigErr struct {
	field string
}

func (err *MissingConfigErr) Error() string {
	return fmt.Sprintf("Missing %s config", err.field)
}

type MailConfig struct {
	host     string
	port     int
	sender   string
	username string
	password string
}

func buildConfig() (*MailConfig, error) {
	host := os.Getenv("MAIL_HOST")

	if host == "" {
		return nil, &MissingConfigErr{"MAIL_HOST"}
	}

	port, portErr := strconv.Atoi(os.Getenv("MAIL_PORT"))

	if portErr != nil {
		return nil, &MissingConfigErr{"MAIL_PORT"}
	}

	sender := os.Getenv("MAIL_SENDER")

	if sender == "" {
		return nil, &MissingConfigErr{"MAIL_SENDER"}
	}

	username := os.Getenv("MAIL_USERNAME")

	if username == "" {
		return nil, &MissingConfigErr{"MAIL_USERNAME"}
	}

	password := os.Getenv("MAIL_PASSWORD")

	if password == "" {
		return nil, &MissingConfigErr{"MAIL_PASSWORD"}
	}

	return &MailConfig{
		host:     host,
		port:     port,
		sender:   sender,
		username: username,
		password: password,
	}, nil
}
