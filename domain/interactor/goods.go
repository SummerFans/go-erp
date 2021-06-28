package interactor

import (
	"go-erp/domain/entity"
	"go-erp/domain/repository"
)

type GoodsInteractor interface {
	Get(id int) (*entity.Goods, error)
}

type goodsInteractor struct {
	goodsRepository   repository.GoodsRepository
	shopifyRepository repository.ShopifyClientRepository
}

func (gi *goodsInteractor) Get(id int) (*entity.Goods, error) {

	return &entity.Goods{
		GoodsID: id,
	}, nil
}

func NewGoodsInteractor(
	goodRepository repository.GoodsRepository,
	shopifyRepository repository.ShopifyClientRepository,
) GoodsInteractor {
	return &goodsInteractor{
		goodsRepository:   goodRepository,
		shopifyRepository: shopifyRepository,
	}
}
