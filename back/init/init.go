package init

import (
	"github.com/PatateDu609/matcha/utils/log"
)

func init() {
	initViper()
	initLogger()
	initDatabase()
	initSession()

	log.Logger.Infoln("everything has been initialized!")
}
