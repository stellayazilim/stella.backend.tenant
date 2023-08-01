package Services

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"log"
)

type ICategoryService interface {
	CreateCategory(data *Types.Category) error
	GetCategoryById(id uint) (Types.Category, error)
	UpdateCategoryById(id uint, data Types.Category) error
	DeleteCategoryById(id uint) error
}

type categoryServcie struct {
}

func CategoryService() ICategoryService {
	return &categoryServcie{}
}

func (s categoryServcie) CreateCategory(data *Types.Category) error {
	if err := DatabaseModule.DB.Create(data).Error; err != nil {
		return fmt.Errorf("Category could not be created")
	}

	return nil
}

func (s categoryServcie) GetCategories(limit int) ([]*Types.Category, error) {
	var categories []*Types.Category
	DatabaseModule.DB.Limit(limit).Preload("Products").Find(&categories)

	return categories, nil
}

func (s categoryServcie) GetCategoryById(id uint) (Types.Category, error) {
	var category Types.Category

	DatabaseModule.DB.Preload("Products").Find(&category, id)

	return category, nil
}

func (s categoryServcie) UpdateCategoryById(id uint, data Types.Category) error {

	data.ID = id
	if err := DatabaseModule.DB.Save(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s categoryServcie) DeleteCategoryById(id uint) error {

	if err := DatabaseModule.DB.Delete(Types.Category{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
