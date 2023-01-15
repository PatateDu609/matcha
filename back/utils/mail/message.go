package mail

import (
	"fmt"
	"strings"

	"github.com/PatateDu609/matcha/config"
)

type Message struct {
	To, Cc, Subject string
	Message         string
}

func (message Message) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("From: %s\r\n", config.Conf.Mail.From))
	builder.WriteString(fmt.Sprintf("To: %s\r\n", message.To))

	if message.Cc != "" {
		builder.WriteString(fmt.Sprintf("Cc: %s\r\n", message.Cc))
	}
	
	builder.WriteString(fmt.Sprintf("Subject: %s\r\n\r\n", message.Subject))

	builder.WriteString(strings.ReplaceAll(message.Message, "\n", "\r\n"))

	return builder.String()
}
