package handler

import (
	"errors"
	"go-erp/domain/interactor"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GoodsHandler struct {
	Handler
	Interactor interactor.GoodsInteractor
}

func (h GoodsHandler) GetGoodsItem(write http.ResponseWriter, request *http.Request, lang, id string) {
	goodsId, err := strconv.Atoi(id)
	if err != nil {
		h.Error(write, errors.New("params error: not goods_id"))
		return
	}
	goods, err := h.Interactor.Get(goodsId)
	if err != nil {
		h.Error(write, errors.New("error"))
		return
	}

	goods.SPU = lang
	goods.CreatedAt = time.Now()
	goods.UpdatedAt = time.Now()
	goods.PublishAt = time.Now()

	successRes := SuccessResponse{
		StatusCode: 200,
		Data:       goods,
	}

	h.Respond(write, 200, successRes)
}

// CreateGood ..
func (h GoodsHandler) CreateGood(c *gin.Context) {
	successRes := SuccessResponse{
		StatusCode: 200,
		Data:       nil,
	}

	h.Respond(c.Writer, 200, successRes)
}
