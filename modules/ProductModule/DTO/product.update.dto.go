package DTO

import (
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/gorm"
)

type ProductUpdateDto struct {
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Explanation *string            `json:"explanation"`
	Sku         *string            `json:"sku"`
	Specs       *map[string]string `json:"specs"`
	Tags        *[]string          `json:"tags"`
	Categories  *[]uint            `json:"categories"`
}

func (d *ProductUpdateDto) ConvertToEntity() Types.Product {
	var categories []*Types.Category
	for _, c := range *d.Categories {
		categories = append(categories, &Types.Category{
			Model: gorm.Model{
				ID: c,
			},
		})
	}
	return Types.Product{
		Name:        *d.Name,
		Description: *d.Description,
		Explanation: *d.Explanation,
		Sku:         *d.Sku,
		Specs:       *d.Specs,
		Tags:        *d.Tags,
		Categories:  categories,
	}
}
