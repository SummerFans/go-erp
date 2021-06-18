package postgres

import (
	"go-erp/domain/entity"
	"go-erp/domain/repository"
	"go-erp/domain/throwable"

	"github.com/go-pg/pg/v10"
)

type goodsRepository struct {
	db *pg.DB
}

func (r *goodsRepository) Get(id int) (*entity.Goods, error) {
	goods := &entity.Goods{GoodsID: id}
	err := r.db.Model().Select(goods)
	if err != nil {
		switch e := err.(type) {
		case pg.Error:
			return nil, e
		default:
			return nil, throwable.NewNotFound("goods not found")
		}
	}
	return goods, nil
}

func NewGoodsRepository(db *pg.DB) repository.GoodsRepository {
	return &goodsRepository{
		db: db,
	}
}
