package config

import (
	"github.com/sirupsen/logrus"
)

const BcryptCost = 12

type Config struct {
	APIPort  int
	Database Database
	LogLevel logrus.Level
}

var Conf Config
