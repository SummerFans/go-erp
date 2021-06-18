package interactor

import (
	"go-erp/domain/entity"
	"go-erp/domain/repository"

	"golang.org/x/text/language"
)

type CategoriesInteractor interface {
	GetAll(params entity.CategoryGetAllParams, locale string) ([]*entity.Category, int8, error)
	Get() (*entity.Category, error)
	Create(category *entity.Category, locales language.Tag) error
	Delete(id int) error
}

type categoriesInteractor struct {
	categoriesRepository repository.CategoryRepository
}

func (gi *categoriesInteractor) GetAll(params entity.CategoryGetAllParams, locale string) ([]*entity.Category, int8, error) {

	categories, count, err := gi.categoriesRepository.GetAll(params, locale)

	return categories, count, err
}

// 获取分类
func (gi *categoriesInteractor) Get() (*entity.Category, error) {

	return &entity.Category{}, nil
}

// 创建分类
func (gi *categoriesInteractor) Create(category *entity.Category, locales language.Tag) error {

	err := gi.categoriesRepository.Create(category, locales)
	if err != nil {
		return err
	}

	return nil
}

func (gi *categoriesInteractor) Delete(id int) error {
	err := gi.categoriesRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoriesInteractor(categoriesRepository repository.CategoryRepository) CategoriesInteractor {
	return &categoriesInteractor{
		categoriesRepository: categoriesRepository,
	}
}
