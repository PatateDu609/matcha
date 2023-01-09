package init

import (
	"github.com/PatateDu609/matcha/config"
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

	conf := viper.New()
	conf.SetDefault("api_port", 4000)
	conf.SetDefault("log_level", "info")
	conf.SetDefault("database", defaultDbConf)

	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(".")
	conf.AddConfigPath("/app")
	if err := conf.ReadInConfig(); err != nil {
		logrus.Warnf("couldn't read config file: %s", err)
	}
	dbConf := defaultDbConf

	if confDatabase, ok := conf.Get("database").(config.Database); ok {
		dbConf = confDatabase
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
		APIPort:  conf.GetInt("api_port"),
		Database: dbConf,
		LogLevel: logrusLevel,
	}
}
