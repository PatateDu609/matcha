package config

import (
	"fmt"
)

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func (conf *Database) DSN() string {
	return fmt.Sprintf("user=%s password=%s sslmode=disable timezone=Europe/Paris host=%s port=%d dbname=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
}
