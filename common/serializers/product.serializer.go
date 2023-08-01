package serializers

import (
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule/DTO"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
)

type IProductSerializer interface {
	SerializeAllFromEntity(product []*Types.Product) []ProductSerializer
	SerializeFromEntity(product *Types.Product) ProductSerializer
	SerializeFromCreateDto(dto *DTO.ProductCreateDto) *Types.Product
	SerializeFromID(id uint) Types.Product
	SerializeAllFromID(dto []uint) []Types.Product
}

type ProductSerializer struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Categories  []CategorySerializer `json:"categories"`
}

func CreateProductSerializer() IProductSerializer {
	return &ProductSerializer{}
}

func (u ProductSerializer) SerializeAllFromEntity(products []*Types.Product) []ProductSerializer {
	var p []ProductSerializer

	for _, product := range products {
		p = append(p, u.SerializeFromEntity(product))
	}
	return p
}

func (u ProductSerializer) SerializeFromEntity(product *Types.Product) ProductSerializer {
	cs := CreateCategorySerializer()
	return ProductSerializer{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Categories:  cs.SerializeAllFromEntity(product.Categories),
	}
}

func (u ProductSerializer) SerializeFromCreateDto(dto *DTO.ProductCreateDto) *Types.Product {
	cs := CreateCategorySerializer()
	return &Types.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Explanation: dto.Explanation,
		Sku:         dto.Sku,
		Specs:       dto.Specs,
		//	Tags:        dto.Tags,
		Categories: cs.SerializeAllFromId(dto.Categories),
	}
}

func (u ProductSerializer) SerializeAllFromID(dto []uint) []Types.Product {
	var products []Types.Product

	for _, product := range dto {
		products = append(products, u.SerializeFromID(product))
	}

	return products
}

func (u ProductSerializer) SerializeFromID(id uint) Types.Product {
	m := gorm.Model{
		ID: id,
	}
	return Types.Product{
		Model: m,
	}
}
