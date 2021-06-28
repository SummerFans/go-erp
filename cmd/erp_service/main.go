package main

import (
	"fmt"
	"go-erp/application/handler"
	"go-erp/dependency"
	"go-erp/domain/interactor"
	"go-erp/presenter"
)

func main() {
	db, err := dependency.NewPostgresConnection()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	} else {
		defer dependency.Close(db)
	}

	shopifyClient, err := dependency.NewShopifyClientConnection()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	categoriesInteractor := interactor.NewCategoriesInteractor(
		dependency.NewCategoryRepository(db),
	)

	goodsInteractor := interactor.NewGoodsInteractor(
		dependency.NewGoodsRepository(db),
		dependency.NewShopifyRepository(shopifyClient),
	)

	categories := handler.CategoriesHandler{Interactor: categoriesInteractor}
	goodsHandler := handler.GoodsHandler{Interactor: goodsInteractor}
	httpServer := presenter.NewHttpServer(
		categories,
		goodsHandler)
	httpServer.Run()
}
