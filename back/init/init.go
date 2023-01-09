package init

import (
	"github.com/PatateDu609/matcha/utils/log"
)

func init() {
	initViper()
	initLogger()
	initDatabase()

	log.Logger.Infoln("everything has been initialized!")
}
