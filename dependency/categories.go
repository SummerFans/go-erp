package dependency

import (
	"go-erp/domain/repository"
	"go-erp/infrastructure/persistence/postgres"

	"github.com/go-pg/pg/v10"
)

func NewCategoryRepository(db interface{}) repository.CategoryRepository {
	switch connection := db.(type) {
	case *pg.DB:
		return postgres.NewCategoriesRepository(connection)
	}
	return nil
}
