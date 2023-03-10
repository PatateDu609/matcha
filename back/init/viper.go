package init

import (
	"github.com/PatateDu609/matcha/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initViper() {
	defaultAPIConf := config.API{
		Host:         "localhost",
		Port:         4000,
		InternalPort: 4001,
	}

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

	defaultSocketIOConf := config.SocketIO{
		URL: "http://localhost:3000",
	}

	defaultSessionConf := config.Session{
		Provider:   "memory",
		Lifetime:   3600,
		CookieName: "gosessid",
	}

	conf := viper.New()
	conf.SetDefault("api", defaultAPIConf)
	conf.SetDefault("log_level", "info")
	conf.SetDefault("database", defaultDbConf)
	conf.SetDefault("redis", defaultRedisConf)
	conf.SetDefault("session", defaultSessionConf)

	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	conf.AddConfigPath("/app")
	if err := conf.ReadInConfig(); err != nil {
		logrus.Warnf("couldn't read config file: %s", err)
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
		API:      defaultAPIConf,
		SocketIO: defaultSocketIOConf,
		Database: defaultDbConf,
		LogLevel: logrusLevel,
		Session:  defaultSessionConf,
	}

	if err := conf.UnmarshalKey("api", &config.Conf.API); err != nil {
		logrus.Errorf("coudln't read API config: %s", err)
	}
	if err := conf.UnmarshalKey("database", &config.Conf.Database); err != nil {
		logrus.Errorf("coudln't read database config: %s", err)
	}
	if err := conf.UnmarshalKey("session", &config.Conf.Session); err != nil {
		logrus.Errorf("coudln't read session config: %s", err)
	}
	if err := conf.UnmarshalKey("oauth", &config.Conf.OAuth); err != nil {
		logrus.Errorf("coudln't read oauth config: %s", err)
	}

	redisConf := defaultRedisConf
	if err := conf.UnmarshalKey("redis", &redisConf); err != nil {
		logrus.Errorf("coudln't read redis config: %s", err)
	}
	config.Conf.RedisClient = redisConf.GetClient()

	if err := conf.UnmarshalKey("socket.io", &config.Conf.SocketIO); err != nil {
		logrus.Errorf("coudln't read redis config: %s", err)
	}

	if err := conf.UnmarshalKey("mail", &config.Conf.Mail); err != nil {
		logrus.Fatalf("couldn't read mail config: %s", err)
	}
	config.Conf.Mail.Authenticate()

	if config.Conf.RedisClient == nil {
		logrus.Fatalf("couldn't connect to redis")
	}
}
