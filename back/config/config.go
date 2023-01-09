package config

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	APIPort  int
	Database Database
	LogLevel logrus.Level
}

var Conf Config
