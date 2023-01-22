package init

import (
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/log"
)

func init() {
	initViper()
	initLogger()
	initDatabase()
	initSession()

	if err := config.Conf.SocketIO.Load(); err != nil {
		log.Logger.Fatalf("invalid socket.io URL `%s`: %s", config.Conf.SocketIO.URL, err)
	}

	log.Logger.Infoln("everything has been initialized!")
}
