package CategoryModule

import (
	"fmt"
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"log"
)

type ICategoryService interface {
	CreateCategory(data *models.Category) error
	GetCategories(limit int) ([]*models.Category, error)
	GetCategoryById(id uint) (models.Category, error)
	UpdateCategoryById(id uint, data models.Category) error
	DeleteCategoryById(id uint) error
}

type categoryServcie struct {
}

func CategoryService() ICategoryService {
	return &categoryServcie{}
}

func (s categoryServcie) CreateCategory(data *models.Category) error {
	if err := DatabaseModule.DB.Create(data).Error; err != nil {
		return fmt.Errorf("Category could not be created")
	}

	return nil
}

func (s categoryServcie) GetCategories(limit int) ([]*models.Category, error) {
	var categories []*models.Category
	DatabaseModule.DB.Limit(limit).Preload("Products").Find(&categories)

	return categories, nil
}

func (s categoryServcie) GetCategoryById(id uint) (models.Category, error) {
	var category models.Category

	DatabaseModule.DB.Preload("Products").Find(&category, id)

	return category, nil
}

func (s categoryServcie) UpdateCategoryById(id uint, data models.Category) error {

	data.ID = id
	if err := DatabaseModule.DB.Save(&data).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (s categoryServcie) DeleteCategoryById(id uint) error {

	if err := DatabaseModule.DB.Delete(models.Category{}, id).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
