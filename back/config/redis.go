package config

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func (conf *Redis) GetClient() *redis.Client {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	}

	return redis.NewClient(opts)
}
