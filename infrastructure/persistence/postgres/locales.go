package postgres

import (
	"go-erp/domain/entity"
	"go-erp/domain/repository"

	"github.com/go-pg/pg/v10"
)

type localeRepository struct {
	db *pg.DB
}

func (r *localeRepository) GetAll() ([]*entity.Locale, error) {

	locale := make([]*entity.Locale, 0)

	return locale, nil
}

func NewLocaleRepository(db *pg.DB) repository.LocaleRepository {
	return &localeRepository{
		db: db,
	}
}
