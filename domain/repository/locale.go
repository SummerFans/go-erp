package repository

import "go-erp/domain/entity"

// GoodsRepository ...
type LocaleRepository interface {
	// 获取语言列表
	GetAll() ([]*entity.Locale, error)
}
