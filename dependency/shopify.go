package dependency

import (
	"fmt"
	"go-erp/conf"
	"go-erp/domain/repository"
	"go-erp/infrastructure"

	goshopify "github.com/bold-commerce/go-shopify"
)

func NewShopifyClientConnection() (*goshopify.Client, error) {
	// 获取Shopify配置
	shopifyClientConf, err := conf.GetShopifyClientConfig()

	if err != nil {
		return nil, fmt.Errorf("shopify API service configuration error: %v", err.Error())
	}

	app := goshopify.App{
		ApiKey:      shopifyClientConf.ApiKey,
		ApiSecret:   shopifyClientConf.ApiSecret,
		RedirectUrl: shopifyClientConf.RedirectUrl,
		Scope:       shopifyClientConf.Scope,
		Password:    shopifyClientConf.Password,
	}

	// 实例化Shopify Client API
	client := goshopify.NewClient(app, shopifyClientConf.Shopname, shopifyClientConf.Token, goshopify.WithVersion(shopifyClientConf.Version))

	_, err = client.Product.Count(nil)
	if err != nil {
		return nil, fmt.Errorf("shopify API service is not available: %v", err.Error())
	}

	return client, nil

}

func NewShopifyRepository(shopfiyClient *goshopify.Client) repository.ShopifyClientRepository {
	return infrastructure.NewShopfiyClientRepository(shopfiyClient)
}
