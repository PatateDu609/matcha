package init

import (
	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/auth/session/storage/memory"
	"github.com/PatateDu609/matcha/auth/session/storage/redis"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/log"
)

func initSession() {
	redis.Register()
	memory.Register()

	globalManager, err := session.NewManager(config.Conf.Session.Provider, config.Conf.Session.CookieName, config.Conf.Session.Lifetime)
	if err != nil {
		log.Logger.Fatalf("couldn't initialize session manager: %v", err)
	}
	session.GlobalSessions = globalManager
}
