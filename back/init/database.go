package init

import (
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initDatabase() {
	conf, err := pgxpool.ParseConfig(config.Conf.Database.DSN())
	if err != nil {
		log.Logger.Fatalf("couldn't parse config for database: %+v", err)
	}

	database.SetupPool(conf)
}
