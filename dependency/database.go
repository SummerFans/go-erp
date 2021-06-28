package dependency

import (
	"fmt"
	"go-erp/conf"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

func NewPostgresConnection() (*pg.DB, error) {
	pgConf, _ := conf.GetDatabaseConfig()

	connection := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", pgConf.HOST, pgConf.PORT),
		User:     pgConf.USER,
		Password: pgConf.PWD,
		Database: pgConf.DB,
	})

	if _, err := connection.Exec("SELECT 1"); err != nil {
		return nil, fmt.Errorf("pg databases service is not available: %v", err.Error())
	}

	return connection, nil
}

func Close(db interface{}) {
	switch connection := db.(type) {
	case *pg.DB:
		connection.Close()
	case *redis.Client:
		connection.Close()
	}
}
