package presenter

import (
	"go-erp/application/handler"

	"github.com/gin-gonic/gin"
)

// goodsRouter ..
func goodsRouter(router *gin.Engine, handler handler.GoodsHandler) {
	goods := router.Group("goods")
	{

		// 获取商品详情
		goods.GET("/:goods_id", func(c *gin.Context) {
			// langStr := getLanguage(c.Request)
			goodsId := c.Param("goods_id")
			handler.GetGoodsItem(c.Writer, c.Request, "", goodsId)
		})

		// 创建商品
		goods.POST("/", func(c *gin.Context) {
			// langStr := getLanguage(c.Request)
			handler.CreateGood(c)
		})
	}

}
