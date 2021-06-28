package conf

import (
	"os"
)

type ShopifyClient struct {
	ApiKey      string
	ApiSecret   string
	RedirectUrl string
	Scope       string // 接口操作范围
	Shopname    string
	Token       string
	Version     string
	Password    string
}

func GetShopifyClientConfig() (ShopifyClient, error) {
	var apiKey, apiSecret, redirectUrl, scope, shopname, token, version, password string
	if apiKey = os.Getenv("SHOPIFY_KEY"); len(apiKey) == 0 {
		apiKey = "be25e1b0b0b8239e0204f1eb46c7d3ad"
		// return ShopifyClient{}, errors.New("SHOPIFY_KEY")
	}
	if apiSecret = os.Getenv("SHOPIFY_SECRET"); len(apiSecret) == 0 {
		apiSecret = "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"
		// return ShopifyClient{}, errors.New("SHOPIFY_SECRET")
	}
	if redirectUrl = os.Getenv("SHOPIFY_REDIRECT_URL"); len(redirectUrl) == 0 {
		redirectUrl = "http://127.0.0.1:4000/callback"
		// return ShopifyClient{}, errors.New("SHOPIFY_REDIRECT_URL")
	}
	if scope = os.Getenv("SHOPIFY_SCOPE"); len(scope) == 0 {
		scope = "read_products,write_products"
		// return ShopifyClient{}, errors.New("SHOPIFY_SCOPE")
	}
	if shopname = os.Getenv("SHOPIFY_SHOPNAME"); len(shopname) == 0 {
		shopname = "dollyface-dev"
		// return ShopifyClient{}, errors.New("SHOPIFY_SHOPNAME")
	}
	if token = os.Getenv("SHOPIFY_TOKEN"); len(token) == 0 {
		token = ""
		// return ShopifyClient{}, errors.New("SHOPIFY_TOKEN")
	}
	if version = os.Getenv("SHOPIFY_VERSION"); len(version) == 0 {
		version = "2021-04"
		// return ShopifyClient{}, errors.New("SHOPIFY_VERSION")
	}
	if password = os.Getenv("SHOPIFY_PASSWORD"); len(password) == 0 {
		password = "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"
		// return ShopifyClient{}, errors.New("SHOPIFY_PASSWORD")
	}

	return ShopifyClient{
		ApiKey:      apiKey,
		ApiSecret:   apiSecret,
		RedirectUrl: redirectUrl,
		Scope:       scope,
		Shopname:    shopname,
		Token:       token,
		Version:     version,
		Password:    password,
	}, nil
}
