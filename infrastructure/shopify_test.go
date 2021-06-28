package infrastructure_test

import (
	"go-erp/dependency"
	"go-erp/infrastructure"
	"testing"
)

func TestShopifySyncGoods(t *testing.T) {
	shopfiyClinent, err := dependency.NewShopifyClientConnection()
	if err != nil {
		t.Errorf("Err: %v", err.Error())
		return
	}
	shopify := infrastructure.NewShopfiyClientRepository(shopfiyClinent)

	if err != nil {
		t.Errorf("Err: %v", err.Error())
		return
	}

	if err = shopify.SyncProduct(); err != nil {
		t.Errorf("Err: %v", err.Error())
		return
	}

	t.Log("Success")
}
