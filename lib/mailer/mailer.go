package mailer

func CreateBuilder() *MailBuilder {
	config, err := buildConfig()
	if err != nil {
		panic(err.Error())
	}

	mail := &MailPayload{
		sender: &config.sender,
	}

	return &MailBuilder{
		config,
		mail,
	}
}
