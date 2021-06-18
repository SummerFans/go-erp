package dependency

import (
	"errors"
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
		return nil, errors.New(fmt.Sprintf("pg databases service is not available: %v", err.Error()))
	}

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// _, err := rdb.Ping(context.Background()).Result()
	// if err != nil {
	// 	return nil, errors.New(fmt.Sprintf("redis databases service is not available: %v", err.Error()))
	// }

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
