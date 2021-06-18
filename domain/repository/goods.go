package repository

import "go-erp/domain/entity"

// GoodsRepository ...
type GoodsRepository interface {
	// 获取获取商品列表
	Get(id int) (*entity.Goods, error)
	// GetAll() ([]*entity.Goods, error)
}
