package postgres

import (
	"fmt"
	"go-erp/domain/entity"
	"go-erp/domain/repository"

	"github.com/go-pg/pg/v10"
	"golang.org/x/text/language"
)

type categoriesRepository struct {
	db *pg.DB
}

func (r *categoriesRepository) Get() (*entity.Category, error) {
	goods := &entity.Category{}
	// err := r.db.Model().Select(goods)
	// if err != nil {
	// 	switch e := err.(type) {
	// 	case pg.Error:
	// 		return nil, e
	// 	default:
	// 		return nil, throwable.NewNotFound("goods not found")
	// 	}
	// }
	return goods, nil
}

func (r *categoriesRepository) GetAll(params entity.CategoryGetAllParams, locale string) ([]*entity.Category, int8, error) {

	if locale == "" {
		locale = "zh_cn"
	}

	categories := make([]*entity.Category, 0)

	sqlStr := fmt.Sprintf("SELECT category.category_id as category_id, code as category_code, category_name, updated_at, deleted_at, created_at FROM \"public\".s_category as category LEFT JOIN \"%s\".s_category_lang as lang ON lang.category_id = category.category_id WHERE deleted_at ISNULL ORDER BY updated_at DESC limit %d offset %d", locale, params.PageSize, params.PageIndex*params.PageSize)

	res, err := r.db.Query(&categories, sqlStr)

	if err != nil {
		return nil, 0, err
	}

	var countNum struct {
		Count int8 `json:"count"`
	}

	_, err = r.db.Query(&countNum, "SELECT count(*) as count FROM \"public\".s_category WHERE deleted_at ISNULL")

	if err != nil {
		return nil, 0, err
	}

	fmt.Println(res.RowsAffected())

	return categories, countNum.Count, nil
}

func (r *categoriesRepository) Create(category *entity.Category, locales language.Tag) error {

	// 创建分类函数
	categorySql := fmt.Sprintf("SELECT create_categories(%d,'%v','%v')", category.ParentID, category.CategoryName, "zh_cn")

	_, err := r.db.Exec(categorySql)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoriesRepository) Delete(id int) error {

	fmt.Println(id)
	_, err := r.db.Exec(`UPDATE s_category SET deleted_at = now() WHERE category_id = ?`, id)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoriesRepository(db *pg.DB) repository.CategoryRepository {
	return &categoriesRepository{
		db: db,
	}
}
