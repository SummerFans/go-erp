package repository

import (
	goshopify "github.com/bold-commerce/go-shopify"
)

type ShopifyClientRepository interface {
	// shopify
	Get(shopifyID int64) (*goshopify.Product, error)

	// 同步Shopify商品
	SyncProduct() error
}
