package serializers

import (
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"gorm.io/gorm"
)

type ICategorySerializer interface {
	SerializeAllFromEntity(categories []*types.Category) []CategorySerializer
	SerializeFromEntity(category *types.Category) CategorySerializer
	SerializeFromCreateDto(dto dto.CategoryCreateDto) types.Category
	SerializeFromID(id uint) *types.Category
	SerializeAllFromId(ids []uint) []*types.Category
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

func (u CategorySerializer) SerializeAllFromEntity(categories []*types.Category) []CategorySerializer {
	var serialized []CategorySerializer

	for _, category := range categories {
		serialized = append(serialized, u.SerializeFromEntity(category))
	}

	return serialized
}

func (u CategorySerializer) SerializeFromEntity(category *types.Category) CategorySerializer {

	ps := CreateProductSerializer()

	return CategorySerializer{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Products:    ps.SerializeAllFromEntity(category.Products),
	}

}

func (u CategorySerializer) SerializeFromCreateDto(dto dto.CategoryCreateDto) types.Category {

	var products []*types.Product
	for _, p := range dto.Products {
		product := types.Product{}
		product.ID = p
		products = append(products, &product)
	}
	return types.Category{
		Name:        dto.Name,
		Description: dto.Description,
		Products:    products,
	}
}
func (u CategorySerializer) SerializeAllFromId(ids []uint) []*types.Category {
	var categories []*types.Category
	for _, ID := range ids {
		categories = append(categories, u.SerializeFromID(ID))
	}
	return categories
}

func (u CategorySerializer) SerializeFromID(id uint) *types.Category {
	m := gorm.Model{
		ID: id,
	}
	return &types.Category{
		Model: m,
	}
}
