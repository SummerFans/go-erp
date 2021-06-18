package presenter

import (
	"go-erp/application/handler"
	"go-erp/application/middleware"

	"github.com/gin-gonic/gin"
)

// goodsRouter ..
func categoryRouter(router *gin.Engine, handler handler.CategoriesHandler) {
	security := middleware.NewSecurityMiddleware()
	router.Use(security.Auth())
	categoriesGroup := router.Group("categories")
	{

		categoriesGroup.GET("/", func(c *gin.Context) {

			handler.GetCategories(c)
		})

		// 创建分类
		categoriesGroup.POST("/", func(c *gin.Context) {
			handler.CreateCategories(c)
		})

		// // 创建商品
		// goods.POST("/", func(c *gin.Context) {
		// 	// langStr := getLanguage(c.Request)
		// 	handler.CreateGood(c)
		// })
	}

}
