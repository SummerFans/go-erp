package handler

import (
	"fmt"
	"go-erp/domain/entity"
	"go-erp/domain/interactor"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type CategoriesHandler struct {
	Handler
	Interactor interactor.CategoriesInteractor
}

// CreateGood ..
func (h CategoriesHandler) GetCategories(c *gin.Context) {

	pageIndex, err := strconv.Atoi(c.DefaultQuery("page_index", "0"))
	if err != nil {
		h.Error(c.Writer, fmt.Errorf("err: page_index is not a numeric type "))
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	if err != nil {
		h.Error(c.Writer, fmt.Errorf("err: page_size is not a numeric type "))
		return
	}

	params := entity.CategoryGetAllParams{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}

	categories, countNum, err := h.Interactor.GetAll(params, "zh_cn")

	if err != nil {
		h.Error(c.Writer, err)
		return

	}
	fmt.Println(countNum)

	successRes := SuccessResponse{
		StatusCode: 0,
		Data:       categories,
	}

	h.Respond(c.Writer, 200, successRes)
}

func (h CategoriesHandler) CreateCategories(c *gin.Context) {

	// localeStr, _ := c.Get("locale")
	// var person Person
	var category entity.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		h.Error(c.Writer, err)
		return
	}

	err := h.Interactor.Create(&category, language.SimplifiedChinese)
	if err != nil {
		h.Error(c.Writer, err)
		return
	}

	successRes := SuccessResponse{
		StatusCode: 0,
		Data:       nil,
	}
	h.Respond(c.Writer, 200, successRes)
}
