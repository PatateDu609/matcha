package config

import (
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

const BcryptCost = 12

type Config struct {
	APIPort     int
	Database    Database
	LogLevel    logrus.Level
	RedisClient *redis.Client
}

var Conf Config
