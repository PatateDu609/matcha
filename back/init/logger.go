package init

import (
	"os"

	"github.com/PatateDu609/matcha/config"
	log2 "github.com/PatateDu609/matcha/utils/log"
	"github.com/sirupsen/logrus"
)

func initLogger() {
	log2.Logger = logrus.New()
	logger := log2.Logger

	logger.SetLevel(config.Conf.LogLevel)
	logger.SetOutput(os.Stderr)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		PadLevelText:           true,
		DisableLevelTruncation: true,
		TimestampFormat:        " 2006-01-02 15:04:05.000 ",
		CallerPrettyfier:       log2.Prettyfier,
		FullTimestamp:          true,
	})
}
