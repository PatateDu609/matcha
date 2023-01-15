package config

import (
	"net/smtp"
)

type Mail struct {
	Address  string `mapstructure:"address"`
	Host     string `mapstructure:"host"`
	From     string `mapstructure:"from"`
	Password string `mapstructure:"password"`
	Identity string `mapstructure:"identity"`

	Auth smtp.Auth
}

func (config *Mail) Authenticate() {
	config.Auth = smtp.PlainAuth(config.Identity, config.From, config.Password, config.Host)
}

func (config *Mail) Send(to []string, message string) error {
	return smtp.SendMail(config.Address, config.Auth, config.From, to, []byte(message))
}
