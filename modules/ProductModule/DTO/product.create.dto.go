package DTO

import (
	"gorm.io/gorm"
)

type ProductCreateDto struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Explanation string            `json:"explanation"`
	Sku         string            `json:"sku"`
	Specs       map[string]string `json:"specs"`
	Tags        []string          `json:"tags"`
	Categories  []uint            `json:"categories"`
}

func (d *ProductCreateDto) ConvertToEntity() types.Product {
	var categories []*types.Category
	for _, c := range d.Categories {
		categories = append(categories, &types.Category{
			Model: gorm.Model{
				ID: c,
			},
		})
	}
	return types.Product{
		Name:        d.Name,
		Description: d.Description,
		Explanation: d.Explanation,
		Sku:         d.Sku,
		Specs:       d.Specs,
		Tags:        d.Tags,
		Categories:  categories,
	}
}
