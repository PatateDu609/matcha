package config

import (
	"fmt"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func (conf *Database) DSN() string {
	return fmt.Sprintf("user=%s password=%s sslmode=disable timezone=Europe/Paris host=%s port=%d dbname=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
}
