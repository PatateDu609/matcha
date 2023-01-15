package mail

import (
	"github.com/PatateDu609/matcha/config"
)

func Send(to, subject string, message string) error {
	msg := Message{
		To:      to,
		Cc:      "",
		Subject: subject,
		Message: message,
	}

	return config.Conf.Mail.Send([]string{msg.To}, msg.String())
}
