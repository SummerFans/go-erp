package infrastructure

import (
	"go-erp/domain/repository"
	"time"

	goshopify "github.com/bold-commerce/go-shopify"
)

type shopifyClientRepository struct {
	shopify *goshopify.Client
}

func (r *shopifyClientRepository) Get(shopifyID int64) (*goshopify.Product, error) {

	product, err := r.shopify.Product.Get(shopifyID, nil)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *shopifyClientRepository) SyncProduct() error {
	nd := time.Now()
	startDate := time.Date(2012, nd.Month(), nd.Day(), nd.Hour(), nd.Minute(), nd.Second(), nd.Nanosecond(), nd.UTC().Location())

	listOptions := &goshopify.ListOptions{
		Limit:        10,
		SinceID:      5600812040357,
		CreatedAtMin: startDate,
	}

	// 1. 获取列表
	_, err := r.shopify.Product.List(listOptions)
	if err != nil {
		return err
	}

	return nil
}

func NewShopfiyClientRepository(shopfiyClient *goshopify.Client) repository.ShopifyClientRepository {
	return &shopifyClientRepository{
		shopify: shopfiyClient,
	}
}
