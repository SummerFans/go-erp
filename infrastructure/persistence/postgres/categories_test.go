package postgres_test

import (
	"go-erp/dependency"
	"go-erp/domain/entity"
	"go-erp/infrastructure/persistence/postgres"
	"testing"

	"golang.org/x/text/language"
)

// TODO 根据列表获取分类
func TestPostgresGetCategories(t *testing.T) {
	db, err := dependency.NewPostgresConnection()
	if err != nil {
		t.Errorf("Err: %v", err.Error())
	}
	categories := postgres.NewCategoriesRepository(db)

	param := entity.CategoryGetAllParams{}
	param.PageSize = 1
	param.PageIndex = 1
	categoriesList, count, err := categories.GetAll(param, "zh_cn")
	if err != nil {
		t.Error(err)
	}
	t.Log(count)

	t.Log(len(categoriesList))
}

// TODO 创建分类
func TestPostgresCreateCategories(t *testing.T) {
	db, err := dependency.NewPostgresConnection()
	if err != nil {
		t.Errorf("Err: %v", err.Error())
	}
	categories := postgres.NewCategoriesRepository(db)

	category := &entity.Category{
		CategoryName: "测试分类2021",
		ParentID:     21,
	}

	err = categories.Create(category, language.SimplifiedChinese)
	if err != nil {
		t.Errorf("Err: %v", err.Error())
		return
	}
	t.Log("Success")

}

// TODO 删除分类
func TestPostgresDeleteCategories(t *testing.T) {
	db, err := dependency.NewPostgresConnection()
	if err != nil {
		t.Errorf("Err: %v", err.Error())
	}
	categories := postgres.NewCategoriesRepository(db)

	err = categories.Delete(30)
	if err != nil {
		t.Errorf("Err: %v", err.Error())
		return
	}

}
