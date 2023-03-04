package config

import (
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

const BcryptCost = 12

type Config struct {
	API         API
	Session     Session
	SocketIO    SocketIO
	Database    Database
	LogLevel    logrus.Level
	RedisClient *redis.Client
	Mail        Mail
	OAuth       OAuth
}

var Conf Config
