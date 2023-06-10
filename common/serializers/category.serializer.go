package serializers

import (
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"github.com/stellayazilim/stella.backend.tenant/models"
	"gorm.io/gorm"
)

type ICategorySerializer interface {
	SerializeAllFromEntity(categories []*models.Category) []CategorySerializer
	SerializeFromEntity(category *models.Category) CategorySerializer
	SerializeFromCreateDto(dto dto.CategoryCreateDto) models.Category
	SerializeFromID(id uint) *models.Category
	SerializeAllFromId(ids []uint) []*models.Category
}

type CategorySerializer struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Products    []ProductSerializer `json:"products"`
}

func CreateCategorySerializer() ICategorySerializer {
	return &CategorySerializer{}
}

func (u CategorySerializer) SerializeAllFromEntity(categories []*models.Category) []CategorySerializer {
	var serialized []CategorySerializer

	for _, category := range categories {
		serialized = append(serialized, u.SerializeFromEntity(category))
	}

	return serialized
}

func (u CategorySerializer) SerializeFromEntity(category *models.Category) CategorySerializer {

	ps := CreateProductSerializer()

	return CategorySerializer{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Products:    ps.SerializeAllFromEntity(category.Products),
	}

}

func (u CategorySerializer) SerializeFromCreateDto(dto dto.CategoryCreateDto) models.Category {

	var products []*models.Product
	for _, p := range dto.Products {
		product := models.Product{}
		product.ID = p
		products = append(products, &product)
	}
	return models.Category{
		Name:        dto.Name,
		Description: dto.Description,
		Products:    products,
	}
}
func (u CategorySerializer) SerializeAllFromId(ids []uint) []*models.Category {
	var categories []*models.Category
	for _, ID := range ids {
		categories = append(categories, u.SerializeFromID(ID))
	}
	return categories
}

func (u CategorySerializer) SerializeFromID(id uint) *models.Category {
	m := gorm.Model{
		ID: id,
	}
	return &models.Category{
		Model: m,
	}
}
