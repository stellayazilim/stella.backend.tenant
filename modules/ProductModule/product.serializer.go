package ProductModule

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
	"time"
)

type IProductSerializer interface {
	SerializeFromEntity(entity *models.Product) productSerializer
	SerializeAllFromEntity(entities []*models.Product) []productSerializer
}

type productCategorySerializer struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type productSerializer struct {
	ID          uint                         `json:"id"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Explanation string                       `json:"explanation"`
	Sku         string                       `json:"sku"`
	Specs       map[string]string            `json:"specs"`
	Tags        []string                     `json:"tags"`
	CreatedAt   time.Time                    `json:"createdAt"`
	UpdatedAt   time.Time                    `json:"updatedAt"`
	Categories  []*productCategorySerializer `json:"categories"`
}

func ProductSerializer() IProductSerializer {
	return &productSerializer{}
}

func (s *productSerializer) SerializeFromEntity(entity *models.Product) productSerializer {
	return productSerializer{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Explanation: entity.Explanation,
		Sku:         entity.Sku,
		Specs:       entity.Specs,
		Tags:        entity.Tags,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Categories: func() []*productCategorySerializer {
			var categories []*productCategorySerializer
			for _, c := range entity.Categories {
				categories = append(categories, &productCategorySerializer{
					ID:          c.ID,
					Name:        c.Name,
					Description: c.Description,
					CreatedAt:   c.CreatedAt,
					UpdatedAt:   c.UpdatedAt,
				})
			}
			return categories
		}(),
	}
}
func (s *productSerializer) SerializeAllFromEntity(entities []*models.Product) []productSerializer {
	var serializer []productSerializer
	for _, c := range entities {
		serializer = append(serializer, s.SerializeFromEntity(c))
	}
	return serializer
}
