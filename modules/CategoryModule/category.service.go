package CategoryModule

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"log"
)

type ICategoryService interface {
	CreateCategory(data *types.Category) error
	GetCategories(limit int) ([]*types.Category, error)
	GetCategoryById(id uint) (types.Category, error)
	UpdateCategoryById(id uint, data types.Category) error
	DeleteCategoryById(id uint) error
}

type categoryServcie struct {
}

func CategoryService() ICategoryService {
	return &categoryServcie{}
}

func (s categoryServcie) CreateCategory(data *types.Category) error {
	if err := DatabaseModule.DB.Create(data).Error; err != nil {
		return fmt.Errorf("Category could not be created")
	}

	return nil
}

func (s categoryServcie) GetCategories(limit int) ([]*types.Category, error) {
	var categories []*types.Category
	DatabaseModule.DB.Limit(limit).Preload("Products").Find(&categories)

	return categories, nil
}

func (s categoryServcie) GetCategoryById(id uint) (types.Category, error) {
	var category types.Category

	DatabaseModule.DB.Preload("Products").Find(&category, id)

	return category, nil
}

func (s categoryServcie) UpdateCategoryById(id uint, data types.Category) error {

	data.ID = id
	if err := DatabaseModule.DB.Save(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s categoryServcie) DeleteCategoryById(id uint) error {

	if err := DatabaseModule.DB.Delete(types.Category{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
