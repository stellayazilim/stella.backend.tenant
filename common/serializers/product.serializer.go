package serializers

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule/DTO"
	"gorm.io/gorm"
)

type IProductSerializer interface {
	SerializeAllFromEntity(product []*models.Product) []ProductSerializer
	SerializeFromEntity(product *models.Product) ProductSerializer
	SerializeFromCreateDto(dto *DTO.ProductCreateDto) *models.Product
	SerializeFromID(id uint) models.Product
	SerializeAllFromID(dto []uint) []models.Product
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

func (u ProductSerializer) SerializeAllFromEntity(products []*models.Product) []ProductSerializer {
	var p []ProductSerializer

	for _, product := range products {
		p = append(p, u.SerializeFromEntity(product))
	}
	return p
}

func (u ProductSerializer) SerializeFromEntity(product *models.Product) ProductSerializer {
	cs := CreateCategorySerializer()
	return ProductSerializer{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Categories:  cs.SerializeAllFromEntity(product.Categories),
	}
}

func (u ProductSerializer) SerializeFromCreateDto(dto *DTO.ProductCreateDto) *models.Product {
	cs := CreateCategorySerializer()
	return &models.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Explanation: dto.Explanation,
		Sku:         dto.Sku,
		Specs:       dto.Specs,
		//	Tags:        dto.Tags,
		Categories: cs.SerializeAllFromId(dto.Categories),
	}
}

func (u ProductSerializer) SerializeAllFromID(dto []uint) []models.Product {
	var products []models.Product

	for _, product := range dto {
		products = append(products, u.SerializeFromID(product))
	}

	return products
}

func (u ProductSerializer) SerializeFromID(id uint) models.Product {
	m := gorm.Model{
		ID: id,
	}
	return models.Product{
		Model: m,
	}
}
