package Types

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Explanation string
	Sku         string
	Specs       map[string]string `gorm:"type:bytea" serializer:"json"`
	Tags        []string          `gorm:"serializer:json"`
	Categories  []*Category       `gorm:"many2many:product_category;"`
	Images      []*Image
}

// single product create request
type ProductCreateRequestBody struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Explanation string            `json:"explanation"`
	Sku         string            `json:"sku"`
	Specs       map[string]string `json:"specs"`
	Tags        []string          `json:"tags"`
	Categories  []*Category       `json:"categories"`
	Images      []*Image          `json:"images"`
}

func (p *ProductCreateRequestBody) ConvertToProduct() *Product {
	return &Product{
		Name:        p.Name,
		Description: p.Description,
		Explanation: p.Explanation,
		Sku:         p.Sku,
		Specs:       p.Specs,
		Tags:        p.Tags,
		Categories:  p.Categories,
		Images:      p.Images,
	}
}

func (p *ProductCreateRequestBody) ConvetToProduct() *Product {
	return &Product{
		Name:        p.Name,
		Description: p.Description,
		Explanation: p.Explanation,
		Sku:         p.Sku,
		Specs:       p.Specs,
		Tags:        p.Tags,
		Categories:  p.Categories,
		Images:      p.Images,
	}
}

// batch product create request
type ProductsCreateRequestBody []*ProductCreateRequestBody

func (p *ProductsCreateRequestBody) ConvertToProductSlice() *[]Product {
	var products *[]Product
	for _, product := range *p {
		*products = append(*products, Product{
			Name:        product.Name,
			Description: product.Description,
			Explanation: product.Explanation,
			Sku:         product.Sku,
			Specs:       product.Specs,
			Tags:        product.Tags,
			Categories:  product.Categories,
			Images:      product.Images})
	}
	return products
}

type ProductSpecsAppendBody map[string]string
type ProductSpecsRemoveBody map[string]string
type ProductSpecsUpdateBody map[string]string

// product update request body
type ProductUpdateReqeustBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Explanation string `json:"explanation"`
	Sku         string `json:"sku"`
}

func (p *ProductUpdateReqeustBody) ConvetToProduct() *Product {
	return &Product{
		Name:        p.Name,
		Description: p.Description,
		Explanation: p.Explanation,
		Sku:         p.Sku,
	}
}
