package serializers

import (
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
)

type ICategorySerializer interface {
	SerializeAllFromEntity(categories []*Types.Category) []CategorySerializer
	SerializeFromEntity(category *Types.Category) CategorySerializer
	SerializeFromCreateDto(dto dto.CategoryCreateDto) Types.Category
	SerializeFromID(id uint) *Types.Category
	SerializeAllFromId(ids []uint) []*Types.Category
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

func (u CategorySerializer) SerializeAllFromEntity(categories []*Types.Category) []CategorySerializer {
	var serialized []CategorySerializer

	for _, category := range categories {
		serialized = append(serialized, u.SerializeFromEntity(category))
	}

	return serialized
}

func (u CategorySerializer) SerializeFromEntity(category *Types.Category) CategorySerializer {

	ps := CreateProductSerializer()

	return CategorySerializer{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Products:    ps.SerializeAllFromEntity(category.Products),
	}

}

func (u CategorySerializer) SerializeFromCreateDto(dto dto.CategoryCreateDto) Types.Category {

	var products []*Types.Product
	for _, p := range dto.Products {
		product := Types.Product{}
		product.ID = p
		products = append(products, &product)
	}
	return Types.Category{
		Name:        dto.Name,
		Description: dto.Description,
		Products:    products,
	}
}
func (u CategorySerializer) SerializeAllFromId(ids []uint) []*Types.Category {
	var categories []*Types.Category
	for _, ID := range ids {
		categories = append(categories, u.SerializeFromID(ID))
	}
	return categories
}

func (u CategorySerializer) SerializeFromID(id uint) *Types.Category {
	m := gorm.Model{
		ID: id,
	}
	return &Types.Category{
		Model: m,
	}
}
