package dependency

import (
	"go-erp/domain/repository"
	"go-erp/infrastructure/persistence/postgres"

	"github.com/go-pg/pg/v10"
)

func NewGoodsRepository(db interface{}) repository.GoodsRepository {
	switch connection := db.(type) {
	case *pg.DB:
		return postgres.NewGoodsRepository(connection)
	}
	return nil
}
