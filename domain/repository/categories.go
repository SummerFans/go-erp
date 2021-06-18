package repository

import (
	"go-erp/domain/entity"

	"golang.org/x/text/language"
)

// GoodsRepository ...
type CategoryRepository interface {
	// 获取分类列表
	GetAll(params entity.CategoryGetAllParams, locale string) ([]*entity.Category, int8, error)
	// 获取分类详情
	Get() (*entity.Category, error)
	// 创建分类
	Create(category *entity.Category, locales language.Tag) error
	// 删除分类
	Delete(id int) error
}
