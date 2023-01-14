package init

import (
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initViper() {
	defaultDbConf := config.Database{
		Host:     "localhost",
		Port:     5432,
		User:     "matcha",
		Password: "password",
		Name:     "matcha",
	}

	defaultRedisConf := config.Redis{
		Host:     "localhost",
		Port:     6379,
		Username: "",
		Password: "",
		DB:       0,
	}

	conf := viper.New()
	conf.SetDefault("api_port", 4000)
	conf.SetDefault("log_level", "info")
	conf.SetDefault("database", defaultDbConf)
	conf.SetDefault("redis", defaultRedisConf)

	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	conf.AddConfigPath("/app")
	if err := conf.ReadInConfig(); err != nil {
		logrus.Warnf("couldn't read config file: %s", err)
	}
	dbConf := defaultDbConf
	redisConf := defaultRedisConf

	if confDatabase, ok := conf.Get("database").(config.Database); ok {
		dbConf = confDatabase
	}
	if confRedis, ok := conf.Get("redis").(config.Redis); ok {
		redisConf = confRedis
	}

	logrusLevel := logrus.InfoLevel
	if levelStr := conf.GetString("log_level"); levelStr != "" {
		if level, err := logrus.ParseLevel(levelStr); err != nil {
			logrus.Warnf("couldn't parse level: %s", err)
		} else {
			logrusLevel = level
		}
	}

	config.Conf = config.Config{
		APIPort:     conf.GetInt("api_port"),
		Database:    dbConf,
		LogLevel:    logrusLevel,
		RedisClient: redisConf.GetClient(),
	}

	if config.Conf.RedisClient == nil {
		log.Logger.Fatalf("couldn't connect to redis")
	}
}
