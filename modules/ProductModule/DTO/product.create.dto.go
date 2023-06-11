package DTO

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
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

func (d *ProductCreateDto) ConvertToEntity() models.Product {
	var categories []*models.Category
	for _, c := range d.Categories {
		categories = append(categories, &models.Category{
			Model: gorm.Model{
				ID: c,
			},
		})
	}
	return models.Product{
		Name:        d.Name,
		Description: d.Description,
		Explanation: d.Explanation,
		Sku:         d.Sku,
		Specs:       d.Specs,
		Tags:        d.Tags,
		Categories:  categories,
	}
}
